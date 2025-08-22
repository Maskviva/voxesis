package ipc

import (
	"fmt"
	"voxesis/src/core/entity"
	"voxesis/src/core/log"
	"voxesis/src/core/process"
	"voxesis/src/services"
)

var (
	mcServer *process.McServer
)

// StartMcServer 启动服务器
func StartMcServer(path string, logCallback func(logLine string)) error {
	if mcServer == nil {
		mcServer = process.NewMcServer()
	}

	log.INFO("22:9", "前端请求启动 MC 服务器，路径:", path)
	// 调用内部的辅助函数来处理复杂逻辑
	return setupAndStartMcServer(path, logCallback)
}

// StopMcServer 停止服务器
func StopMcServer() error {
	log.INFO("29:9", "前端请求停止 MC 服务器...")
	// 将请求委托给 mcServer 实例
	return mcServer.Stop()
}

// SendCommandToMcServer 发送命令
func SendCommandToMcServer(command string) error {
	log.INFO("36:9", "前端请求向 MC 服务器发送命令: '%s'", command)
	// 将请求委托给 mcServer 实例
	return mcServer.SendCommand(command)
}

// GetMcServerStatus 获取状态
func GetMcServerStatus() (entity.McServerState, error) {
	var state entity.McServerState

	if mcServer == nil {
		return state, fmt.Errorf("进程未启动")
	}

	serverStatus := services.GetMinecraftServerStatus("127.0.0.1", 19132)
	processState, err := mcServer.GetStatus()

	if err != nil {
		return state, err
	}

	// 安全地处理可能为nil的指针
	if serverStatus.MOTD != nil {
		state.MOTD = *serverStatus.MOTD
	}

	if serverStatus.Protocol != nil {
		state.Protocol = fmt.Sprintf("%d", *serverStatus.Protocol)
	}

	if serverStatus.Version != nil {
		state.Version = *serverStatus.Version
	}

	if serverStatus.PlayersOnline != nil {
		state.PlayersOnline = fmt.Sprintf("%d", *serverStatus.PlayersOnline)
	}

	if serverStatus.PlayersMax != nil {
		state.PlayersMax = fmt.Sprintf("%d", *serverStatus.PlayersMax)
	}

	if serverStatus.ServerID != nil {
		state.ServerID = *serverStatus.ServerID
	}

	if serverStatus.LevelName != nil {
		state.LevelName = *serverStatus.LevelName
	}

	if serverStatus.GameModeID != nil {
		state.GameModeID = *serverStatus.GameModeID
	}

	if serverStatus.PortV4 != nil {
		state.PortV4 = fmt.Sprintf("%d", *serverStatus.PortV4)
	}

	if serverStatus.PortV6 != nil {
		state.PortV6 = fmt.Sprintf("%d", *serverStatus.PortV6)
	}

	state.RunTime = processState.RunTime
	state.Memory = processState.Memory
	state.Cpu = processState.Cpu
	state.Pid = processState.Pid

	// 将请求委托给 mcServer 实例
	return state, nil
}

// setupAndStartMcServer 是一个内部函数，它包含了创建闭包的复杂逻辑
func setupAndStartMcServer(path string, logCallback func(logLine string)) error {
	log.INFO("115:9", "即将调用 mcServer.Start ...")
	err := mcServer.Start(path, logCallback)
	if err != nil {
		log.ERROR("109:21", "mcServer.Start 调用返回错误: %v", err)
	} else {
		log.INFO("109:21", "mcServer.Start 调用成功返回。")
	}

	return err
}
