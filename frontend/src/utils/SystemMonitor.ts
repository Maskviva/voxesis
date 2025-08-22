import {Quit, WindowIsMaximised, WindowMaximise, WindowMinimise, WindowUnmaximise} from '../../wailsjs/runtime'
import {
    GetAllAppConfig,
    GetAppConfigByKey,
    GetMcAllConfig,
    GetMcConfigByKey,
    OpenMcServerDirectoryDialog,
    OpenProxyServerDirectoryDialog,
    UpDataAppConfig,
    UpDataMcConfig
} from "../../wailsjs/go/src/App"
import {entity} from "../../wailsjs/go/models"

export function winMinimize() {
    WindowMinimise();
}

export function winMaximize(value: boolean) {
    value ? WindowMaximise() : WindowUnmaximise();
}

export function closeWin() {
    // Wails 没有直接提供关闭窗口的方法，但可以使用 Quit 退出整个应用
    // 或者根据需要隐藏窗口
    // WindowHide();
    // 如果确实需要关闭窗口，可以使用 Quit
    Quit();
}

// 窗口变化监听
export function watchWindowState(callback: (isMaximized: boolean) => void) {
    // Wails 没有直接提供窗口大小变化和移动的监听器
    // 但我们可以通过轮询方式检查窗口状态
    const checkWindowState = async () => {
        try {
            const isMaximized = await WindowIsMaximised();
            callback(isMaximized);
        } catch (error) {
            console.error('检查窗口状态时出错:', error);
        }
    };

    // 定期检查窗口状态（例如每500毫秒）
    const intervalId = setInterval(checkWindowState, 500);

    // 返回清理函数
    return () => {
        clearInterval(intervalId);
    };
}

export function getAllAppConfig(): Promise<entity.AppConfig> {
    return GetAllAppConfig();
}

export function getAppConfigByKey(key: string): Promise<string> {
    return GetAppConfigByKey(key);
}

export function upDataAppConfig(key: string, value: any): Promise<void> {
    return UpDataAppConfig(key, value);
}

export function getMcAllConfig(): Promise<entity.McServerConfig> {
    return GetMcAllConfig();
}

export function getMcConfigByKey(key: string): Promise<string> {
    return GetMcConfigByKey(key);
}

export function upDataMcConfig(key: string, value: string): Promise<void> {
    return UpDataMcConfig(key, value);
}

export function openMcServerDirectoryDialog(): Promise<string> {
    return OpenMcServerDirectoryDialog()
}

export function openProxyServerDirectoryDialog(): Promise<string> {
    return OpenProxyServerDirectoryDialog()
}