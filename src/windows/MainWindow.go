package windows

import (
	"embed"
	"voxesis/src"
	"voxesis/src/core/log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

func NewMainWindow(assets embed.FS, app *src.App) {
	err := wails.Run(&options.App{
		Title:     "Voxesis",
		Width:     1024,
		Height:    768,
		MinWidth:  580,
		MinHeight: 540,
		Frameless: true, // 隐藏窗口标题栏
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		log.ERROR("MainWindow.go 14:18", err.Error())
	}
}
