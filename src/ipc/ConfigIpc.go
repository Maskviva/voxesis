package ipc

import (
	"voxesis/src/core/config"
	ImplConfig "voxesis/src/core/config/impl"
	"voxesis/src/core/entity"
)

var (
	ConfigManager config.Manager
)

func InitConfigIpc(appDir string) error {
	ConfigManager = ImplConfig.NewManager()

	err := ConfigManager.LoadConfig(appDir+"/config/config.json", "E:/MCBESERVER2")
	return err
}

func GetAllAppConfig() (*entity.AppConfig, error) {
	appConfig, err := ConfigManager.GetAllAppConfig()

	return appConfig, err
}

func GetAppConfigByKey(key string) (string, error) {
	value, err := ConfigManager.GetAppConfigByKey(key)

	return value, err
}

func UpDataAppConfig(key string, value interface{}) error {
	err := ConfigManager.UpDataAppConfig(key, value)

	return err
}

func GetMcAllConfig() (*entity.McServerConfig, error) {
	mcConfig, err := ConfigManager.GetMcAllConfig()

	return mcConfig, err
}

func GetMcConfigByKey(key string) (string, error) {
	value, err := ConfigManager.GetMcConfigByKey(key)

	return value, err
}

func UpDataMcConfig(key string, value string) error {
	err := ConfigManager.UpDataMcConfig(key, value)

	return err
}
