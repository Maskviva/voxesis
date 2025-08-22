package main

import (
	"embed"
	"voxesis/src/windows"

	"voxesis/src"
)

var (
	//go:embed all:frontend/dist
	assets embed.FS

	//go:embed build/windows/icon.ico
	icoIcon []byte

	//go:embed build/windows/icon-systray.ico
	systrayIcon []byte

	app *src.App
)

func main() {

	// 创建app
	app = src.NewApp()

	// 启动系统托盘
	go windows.NewMainSystray(systrayIcon)

	// 创建窗口
	windows.NewMainWindow(assets, app)
}
