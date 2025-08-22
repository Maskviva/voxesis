package ipc

import (
	"context"
	_ "embed"
	"fmt"
	"voxesis/src/core/utils"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	mcNativeServerFileList = []string{
		"bedrock_server.exe",
		"server.properties",
	}

	mcLeviLaminaServerFileList = []string{
		"bedrock_server_mod.exe",
		"server.properties",
	}

	proxyServerFileList = []string{
		"bedrock-server-proxy.exe",
	}
)

func OpenMcServerDirectoryDialog(ctx context.Context) (string, error) {
	dialog, err := utils.OpenDirectoryDialog(ctx, "选择Minecraft服务器目录", "C:/")
	if err != nil {
		return "", err
	}

	err = utils.AuthDirectory(dialog, mcNativeServerFileList, mcLeviLaminaServerFileList)
	if err != nil && err.Error() != "open : The system cannot find the file specified." {
		var feedback string

		if err.Error() == "无法使用的目录" {
			feedback = "请选择正确的MC服务器目录"
		} else {
			feedback = fmt.Sprintf("发生错误: %s", err.Error())
		}

		options := runtime.MessageDialogOptions{
			Type:          runtime.ErrorDialog,
			Title:         "Voxesis-目录选择器",
			Message:       feedback,
			Buttons:       []string{"确定"},
			DefaultButton: "确定",
		}

		_, err := runtime.MessageDialog(ctx, options)

		return "", err
	}
	return dialog, nil
}

func OpenProxyServerDirectoryDialog(ctx context.Context) (string, error) {
	dialog, err := utils.OpenDirectoryDialog(ctx, "选择代理服务器目录", "C:/")
	if err != nil {
		return "", err
	}

	err = utils.AuthDirectory(dialog, proxyServerFileList, proxyServerFileList)
	if err != nil && err.Error() != "open : The system cannot find the file specified." {
		var feedback string

		if err.Error() == "无法使用的目录" {
			feedback = "请选择正确的代理服务器目录"
		} else {
			feedback = fmt.Sprintf("发生错误: %s", err.Error())
		}

		options := runtime.MessageDialogOptions{
			Type:          runtime.ErrorDialog,
			Title:         "Voxesis-目录选择器",
			Message:       feedback,
			Buttons:       []string{"确定"},
			DefaultButton: "确定",
		}

		_, err := runtime.MessageDialog(ctx, options)

		return "", err
	}

	return dialog, nil
}
