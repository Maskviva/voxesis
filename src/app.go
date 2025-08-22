package src

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"voxesis/src/core/entity"
	"voxesis/src/core/log"
	"voxesis/src/ipc"
	"voxesis/src/services"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp 创建一个新的应用程序应用程序结构
func NewApp() *App {
	return &App{}
}

// Startup 在应用程序启动时调用。保存上下文
// 因此，我们可以调用运行时方法
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx

	// 初始化配置管理器
	appDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("获取工作目录失败: %v\n", err)
		return
	}

	err = log.InitLogger(filepath.Join(appDir, "log"), "app.log")
	if err != nil {
		fmt.Printf("日志系统初始化失败: %v\n", err)
		return
	}

	// 确保配置文件存在
	configFilePath := filepath.Join(appDir, "config", "config.json")
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		// 确保config目录存在
		configDir := filepath.Join(appDir, "config")

		if err := os.MkdirAll(configDir, 0755); err != nil {
			log.ERROR("app.go 49:22", "创建config目录失败: ", err.Error())
		} else {
			// 创建配置文件并写入初始内容
			file, err := os.Create(configFilePath)

			if err != nil {
				log.ERROR("app.go 53:29", "创建config文件失败: ", err.Error())
			} else {

				// 写入空的JSON对象作为初始内容
				_, err = file.WriteString("{}")
				if err != nil {
					log.ERROR("app.go 58:31", "初始化config文件失败: ", err.Error())
				}

				// 关闭文件
				err = file.Close()
				if err != nil {
					log.ERROR("app.go 63:28", err.Error())
				}
			}
		}
	} else if err != nil {
		// Stat函数出现其他错误
		log.ERROR("app.go 46:21", "检查config文件时出错: ", err.Error())
	}

	// 初始化IPC配置
	err = ipc.InitConfigIpc(appDir)
	if err != nil {
		log.ERROR("app.go 75:15", "初始化配置管理器失败: ", err.Error())
	}
}

// GetAllAppConfig 获取所有应用程序配置
func (a *App) GetAllAppConfig() *entity.AppConfig {
	appConfig, err := ipc.GetAllAppConfig()

	if err != nil {
		log.ERROR("app.go 83:27", "获取应用程序配置失败: ", err.Error())
	}

	return appConfig
}

// GetAppConfigByKey 获取应用程序配置
func (a *App) GetAppConfigByKey(key string) string {
	value, err := ipc.GetAppConfigByKey(key)
	if err != nil {
		log.ERROR("app.go 89:27", "获取应用程序配置失败: ", err.Error())
	}
	return value
}

// UpDataAppConfig 更新应用程序配置
func (a *App) UpDataAppConfig(key string, value interface{}) error {
	err := ipc.UpDataAppConfig(key, value)

	if err != nil {
		log.ERROR("app.go 103:16", "更新应用程序配置失败: ", err.Error())
		return err
	}

	return nil
}

// GetMcAllConfig 获取所有Minecraft服务器配置
func (a *App) GetMcAllConfig() *entity.McServerConfig {
	config, err := ipc.GetMcAllConfig()

	if err != nil {
		log.ERROR("app.go 112:24", "获取所有Minecraft服务器配置失败: ", err)
	}

	return config
}

// GetMcConfigByKey 获取Minecraft服务器配置
func (a *App) GetMcConfigByKey(key string) string {
	value, err := ipc.GetMcConfigByKey(key)

	if err != nil {
		log.ERROR("app.go 123:23", "获取Minecraft服务器配置失败: ", err.Error())
	}

	return value
}

// UpDataMcConfig 更新Minecraft服务器配置
func (a *App) UpDataMcConfig(key string, value string) {
	err := ipc.UpDataMcConfig(key, value)

	if err != nil {
		log.ERROR("app.go 134:16", "更新Minecraft服务器配置失败: ", err.Error())
	}
}

// OpenMcServerDirectoryDialog 打开Minecraft服务器目录对话框
func (a *App) OpenMcServerDirectoryDialog() string {
	dialog, err := ipc.OpenMcServerDirectoryDialog(a.ctx)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	if dialog == "" {
		return ""
	}

	err = a.UpDataAppConfig("mc_server_root_path", dialog)
	if err != nil {
		log.ERROR("app.go 157:20", err.Error())
		return ""
	}

	return dialog
}

// OpenProxyServerDirectoryDialog 打开代理服务器目录对话框
func (a *App) OpenProxyServerDirectoryDialog() string {
	dialog, err := ipc.OpenProxyServerDirectoryDialog(a.ctx)
	if err != nil {
		log.ERROR("app.go 167:20", err.Error())
		return ""
	}

	if dialog == "" {
		return ""
	}

	err = a.UpDataAppConfig("mc_server_proxy_root_path", dialog)
	if err != nil {
		log.ERROR("app.go 177:13", err.Error())
		return ""
	}

	return dialog
}

// StartMcServer 是暴露给前端的、接口干净的函数
func (a *App) StartMcServer() string {
	logCallback := func(logLine string) {
		if a.ctx != nil {
			runtime.EventsEmit(a.ctx, "mc_server_log", logLine)
		}
	}

	path, err := ipc.ConfigManager.GetAppConfigByKey("mc_server_root_path")

	if err != nil {
		log.ERROR("app.go 198:36", err.Error())
		return ""
	}

	// 验证文件是否存在
	if _, err := os.Stat(path + "\\bedrock_server.exe"); err == nil {
		path = path + "\\bedrock_server.exe"
	} else if _, err := os.Stat(path + "\\bedrock_server_mod.exe"); err == nil {
		path = path + "\\bedrock_server_mod.exe"
	} else {
		log.ERROR("211:13", "无法找到 bedrock_server.exe 或 bedrock_server_mod.exe 文件")
		return ""
	}

	err = ipc.StartMcServer(path, logCallback)

	if err != nil {
		log.ERROR("app.go 218:43", err.Error())
		return ""
	}

	return path
}

// StopMcServer 暴露给前端，用于停止服务器
func (a *App) StopMcServer() {
	err := ipc.StopMcServer()

	if err != nil {
		log.ERROR("app.go 230:16", err.Error())
	}
}

// SendCommandToMcServer 暴露给前端，用于发送命令
func (a *App) SendCommandToMcServer(command string) {
	err := ipc.SendCommandToMcServer(command)

	if err != nil {
		log.ERROR("app.go 239:16", err.Error())
	}
}

// GetMcServerStatus 暴露给前端，用于获取状态
func (a *App) GetMcServerStatus() (entity.McServerState, error) {
	state, err := ipc.GetMcServerStatus()

	if err != nil && err.Error() != "服务器未在运行" && err.Error() != "进程未在启动" {
		log.ERROR("app.go 245:23", err.Error())
	}

	return state, err
}

func (a *App) GetOsState() *entity.OsState {
	state, err := services.GetOsStateListener()

	if err != nil {
		log.ERROR("app.go 258:28", err.Error())
	}

	return state
}

func (a *App) SendMessageToLLOneBot(message string) {
	err := services.SendMessageToLLOneBot(ipc.ConfigManager, message)

	if err != nil {
		log.ERROR("app.go 268:21", err.Error())
		return
	}
}
