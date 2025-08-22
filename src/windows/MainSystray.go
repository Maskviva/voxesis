package windows

import (
	"github.com/getlantern/systray"
)

var Icon []byte

func NewMainSystray(icoIcon []byte) {
	Icon = icoIcon

	systray.Run(onReady, onExit)
}

// onReady 是系统托盘初始化函数，在应用启动时被调用
func onReady() {
	systray.SetIcon(Icon)
	systray.SetTitle("Voxesis")
	systray.SetTooltip("Voxesis 控制台")

	mShow := systray.AddMenuItem("显示", "显示窗口")
	mQuit := systray.AddMenuItem("退出", "退出应用")

	go func() {
		for {
			select {
			case <-mShow.ClickedCh:
				// 显示窗口的逻辑将在后续实现
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()
}

func onExit() {
	// 清理工作
}
