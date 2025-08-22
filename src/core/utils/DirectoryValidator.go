package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func OpenDirectoryDialog(ctx context.Context, title string, defaultPath string) (string, error) {
	options := runtime.OpenDialogOptions{
		Title:                title,
		DefaultDirectory:     defaultPath,
		ShowHiddenFiles:      false,
		CanCreateDirectories: true,
		ResolvesAliases:      true,
	}

	return runtime.OpenDirectoryDialog(ctx, options)
}

func AuthDirectory(path string, fileList []string, file2List []string) error {
	dir, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func(dir *os.File) {
		_ = dir.Close()
	}(dir)

	files, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	checkFiles := func(list []string) error {
		requiredFiles := make(map[string]bool)
		for _, file := range list {
			requiredFiles[file] = true
		}

		fileMap := make(map[string]bool)
		for _, file := range files {
			if !file.IsDir() {
				fileMap[file.Name()] = true
			}
		}

		for file := range requiredFiles {
			if !fileMap[file] {
				return fmt.Errorf("无法使用的目录")
			}
		}
		return nil
	}

	err1 := checkFiles(fileList)
	if err1 == nil {
		return nil
	}

	err2 := checkFiles(file2List)
	if err2 == nil {
		return nil
	}

	return fmt.Errorf("无法使用的目录")
}
