package services

import (
	"encoding/json"
	"net/http"
	"strings"
	"voxesis/src/core/config"
)

func SendMessageToLLOneBot(ConfigManager config.Manager, message string) error {
	Token, err := ConfigManager.GetAppConfigByKey("llonebot_token")
	Port, err := ConfigManager.GetAppConfigByKey("qq_bot_port")
	group, err := ConfigManager.GetAppConfigByKey("qq_group")

	if err != nil {
		return err
	}

	client := &http.Client{}

	// 创建一个map来存储JSON数据
	data := map[string]interface{}{
		"group_id":     group,
		"message_type": "group",
		"message":      message,
	}

	// 使用json.Marshal正确地序列化JSON数据
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// 发送JSON数据
	req, err := http.NewRequest("POST", "http://127.0.0.1:"+Port+"/send_group_msg", strings.NewReader(string(jsonData)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", Token)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	_ = resp.Body.Close()

	return nil
}
