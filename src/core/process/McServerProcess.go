package process

import (
	"fmt"
	"path/filepath"
	"voxesis/src/core/entity"
	BaseProcess "voxesis/src/core/process/base"
)

// McServer 结构体负责管理一个 Minecraft 服务器进程的完整生命周期。
// 它封装了一个通用的 ConPtyProcessManager。
type McServer struct {
	manager *BaseProcess.ConPtyProcessManager
}

// NewMcServer 创建一个新的 McServer 管理器实例。
func NewMcServer() *McServer {
	return &McServer{}
}

// Start 启动服务器进程。它负责创建底层的 ProcessManager。
// path: 可执行文件的完整路径。
// logCallback: 用于处理日志输出的回调函数。
func (s *McServer) Start(path string, logCallback func(log string)) error {
	// 如果已有实例在运行，先停止
	if s.manager != nil && s.manager.IsRunning() {
		if err := s.manager.Stop(); err != nil {
			return fmt.Errorf("无法停止正在运行的旧服务器: %w", err)
		}
	}

	// 创建一个新的底层管理器
	var err error
	if s.manager, err = BaseProcess.NewConPtyProcessManager(path); err != nil {
		return err
	}

	// 设置回调，包装原始回调以移除 ANSI 转义字符
	s.manager.SetOutputCallback(func(log string) {
		logCallback(log)
	})

	// 启动进程
	workingDir := filepath.Dir(path)
	return s.manager.Start(workingDir)
}

// Stop 停止服务器进程。
func (s *McServer) Stop() error {
	if s.manager == nil || !s.manager.IsRunning() {
		return nil // 未运行，视为成功停止
	}
	return s.manager.Stop()
}

// SendCommand 向服务器发送命令。
func (s *McServer) SendCommand(command string) error {
	if s.manager == nil || !s.manager.IsRunning() {
		return fmt.Errorf("服务器未在运行")
	}
	return s.manager.SendCommand(command)
}

// IsRunning 检查服务器是否在运行。
func (s *McServer) IsRunning() bool {
	if s.manager == nil {
		return false
	}
	return s.manager.IsRunning()
}

// GetStatus 获取服务器进程的状态。
func (s *McServer) GetStatus() (entity.ProcessState, error) {
	if s.manager == nil || !s.manager.IsRunning() {
		return entity.ProcessState{}, fmt.Errorf("服务器未在运行")
	}
	return s.manager.GetProcessStatus()
}
