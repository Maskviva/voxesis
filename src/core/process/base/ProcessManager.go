package BaseProcess

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"sync"
	"syscall"
	"time"
	"voxesis/src/core/entity"

	"github.com/shirou/gopsutil/v3/process"
)

// ProcessManager 负责管理一个独立的外部进程。
// 此版本使用 WinPTY 方案以兼容需要真实终端环境的程序。
// 此结构体的所有方法都是并发安全的。
type ProcessManager struct {
	mu             sync.RWMutex
	binary         string
	cmd            *exec.Cmd
	stdin          io.WriteCloser // 使用标准输入管道与 WinPTY 通信
	proc           *process.Process
	outputCallback func(log string)

	// 用于准确监控 CPU 的字段
	cpuPercentCache float64
	stopMonitorChan chan struct{}
}

// NewProcessManager 为给定的可执行文件路径创建一个新的进程管理器。
func NewProcessManager(path string) (*ProcessManager, error) {
	binaryPath, err := exec.LookPath(path)
	if err != nil {
		return nil, fmt.Errorf("未能找到可执行文件 '%s': %w", path, err)
	}
	return &ProcessManager{binary: binaryPath}, nil
}

// SetOutputCallback 设置一个回调函数来处理进程的标准输出和标准错误。
func (pm *ProcessManager) SetOutputCallback(callback func(log string)) {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.outputCallback = callback
}

// Start 使用给定的命令行参数和工作目录来执行进程，并启动后台监控。
// 【核心实现】: 通过 winpty.exe 启动目标进程，并使用标准管道进行通信。
func (pm *ProcessManager) Start(workingDir string, options ...string) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if pm.proc != nil {
		if isRunning, _ := pm.proc.IsRunning(); isRunning {
			return fmt.Errorf("此管理器已在运行一个进程，PID 为 %d", pm.proc.Pid)
		}
	}

	// WinPTY 方案仅在 Windows 上有效和必要
	if runtime.GOOS != "windows" {
		return fmt.Errorf("WinPTY 方案仅支持 Windows 平台")
	}

	// 构造被 WinPTY 调用的完整命令，格式为：[a.binary, option1, option2, ...]
	targetCmd := append([]string{pm.binary}, options...)

	// 最终执行的命令是 "winpty.exe a.binary option1 option2..."
	// 假设 winpty.exe 就在当前目录下或在系统 PATH 中
	pm.cmd = exec.Command("winpty.exe", targetCmd...)
	pm.cmd.Dir = workingDir
	pm.cmd.Env = os.Environ()
	pm.cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	// 使用标准管道连接到 winpty.exe 的输入输出
	stdoutPipe, err := pm.cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("创建标准输出管道失败: %w", err)
	}
	stderrPipe, err := pm.cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("创建标准错误管道失败: %w", err)
	}
	pm.stdin, err = pm.cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("创建标准输入管道失败: %w", err)
	}

	if pm.outputCallback != nil {
		go pm.readPipe(stdoutPipe, "[STDOUT]")
		go pm.readPipe(stderrPipe, "[STDERR]") // winpty 可能将所有输出都重定向到 stdout，但保留以防万一
	}

	if err := pm.cmd.Start(); err != nil {
		return fmt.Errorf("启动 winpty 进程失败: %w", err)
	}

	go func() { _ = pm.cmd.Wait() }()

	// 监控的是 winpty.exe 进程。当它结束时，其子进程 (BDS) 也会结束。
	pm.proc, err = process.NewProcess(int32(pm.cmd.Process.Pid))
	if err != nil {
		_ = pm.cmd.Process.Kill()
		return fmt.Errorf("进程已启动但创建监控器失败: %w", err)
	}

	pm.stopMonitorChan = make(chan struct{})
	go pm.monitorCPU()

	log.Printf("进程已通过 WinPTY (管道模式) 启动, PID: %d。", pm.cmd.Process.Pid)
	return nil
}

// monitorCPU 是一个后台 goroutine，定期、准确地更新 CPU 使用率。
func (pm *ProcessManager) monitorCPU() {
	_, _ = pm.proc.CPUPercent() // 第一次调用用于初始化

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-pm.stopMonitorChan:
			log.Printf("PID %d 的 CPU 监控已停止。", pm.proc.Pid)
			return
		case <-ticker.C:
			percent, err := pm.proc.CPUPercent()
			if err != nil {
				return // 进程已退出，停止监控
			}

			pm.mu.Lock()
			pm.cpuPercentCache = percent
			pm.mu.Unlock()
		}
	}
}

// Stop 优雅地终止进程，并停止后台监控。
func (pm *ProcessManager) Stop() error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if pm.stopMonitorChan != nil {
		close(pm.stopMonitorChan)
		pm.stopMonitorChan = nil
	}

	if pm.proc == nil {
		return nil
	}
	if isRunning, _ := pm.proc.IsRunning(); !isRunning {
		pm.resetState()
		return nil
	}

	// 终止 winpty.exe 进程，它会负责清理其子进程
	if err := pm.cmd.Process.Kill(); err != nil {
		log.Printf("强制杀死进程 PID %d 失败: %v", pm.proc.Pid, err)
	}

	// 等待 Wait 完成并清理资源
	<-time.After(1 * time.Second) // 给予一点时间让系统清理
	pm.resetState()
	return nil
}

// GetProcessStatus 返回进程的当前资源使用情况，CPU部分从缓存读取。
func (pm *ProcessManager) GetProcessStatus() (entity.ProcessState, error) {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	state := entity.ProcessState{}

	if pm.proc == nil {
		return state, nil
	}

	isRunning, err := pm.proc.IsRunning()
	if err != nil || !isRunning {
		return state, nil
	}

	state.Pid = fmt.Sprintf("%d", pm.proc.Pid)

	// CPU 校准逻辑 (原样保留)
	cpuCores := runtime.NumCPU()
	var taskManagerCpuPercent float64

	if cpuCores > 0 {
		percentOfTotal := pm.cpuPercentCache / float64(cpuCores)
		taskManagerCpuPercent = percentOfTotal * 2
	} else {
		taskManagerCpuPercent = pm.cpuPercentCache
	}
	if taskManagerCpuPercent > 100.0 {
		taskManagerCpuPercent = 100.0
	}
	state.Cpu = taskManagerCpuPercent

	if memInfo, err := pm.proc.MemoryInfo(); err == nil {
		memoryMB := float64(memInfo.RSS) / 1024 / 1024
		state.Memory = memoryMB
	} else {
		state.Memory = 0
	}

	if createTimeMs, err := pm.proc.CreateTime(); err == nil {
		createTime := time.Unix(0, createTimeMs*int64(time.Millisecond))
		uptime := time.Since(createTime).Round(time.Second)
		state.RunTime = uptime.String()
	} else {
		state.RunTime = "unknow"
	}

	return state, nil
}

// readPipe 从管道中读取数据并调用回调函数。
func (pm *ProcessManager) readPipe(pipe io.ReadCloser, prefix string) {
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		line := scanner.Text()
		pm.mu.RLock()
		if pm.outputCallback != nil {
			pm.outputCallback(fmt.Sprintf("%s %s", prefix, line))
		}
		pm.mu.RUnlock()
	}
}

// resetState 清理进程状态。
func (pm *ProcessManager) resetState() {
	pm.cmd = nil
	pm.proc = nil
	pm.stdin = nil
	pm.cpuPercentCache = 0
}

// IsRunning 检查被管理的进程当前是否正在运行。
func (pm *ProcessManager) IsRunning() bool {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	if pm.proc == nil {
		return false
	}
	running, err := pm.proc.IsRunning()
	return err == nil && running
}

// SendCommand 通过标准输入管道向 WinPTY 发送命令。
func (pm *ProcessManager) SendCommand(command string) error {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	if pm.stdin == nil {
		return fmt.Errorf("进程未启动或标准输入不可用")
	}
	if len(command) == 0 || command[len(command)-1] != '\n' {
		command += "\n"
	}

	// 直接写入 UTF-8 编码的命令。WinPTY 会模拟键盘输入到 BDS 终端中，确保命令被正确解析。
	if _, err := pm.stdin.Write([]byte(command)); err != nil {
		return fmt.Errorf("向 WinPTY 写入命令失败: %w", err)
	}
	return nil
}
