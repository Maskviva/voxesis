package services

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v3/process"
)

type ProcessContainer struct {
	mu      sync.Mutex
	process *process.Process
}

// MAGIC Minecraft 基岩版协议常量
var MAGIC = []byte{
	0x00, 0xFF, 0xFF, 0x00, 0xFE, 0xFE, 0xFE, 0xFE,
	0xFD, 0xFD, 0xFD, 0xFD, 0x12, 0x34, 0x56, 0x78,
}

// ServerStatus 表示服务器状态的响应结构。
type ServerStatus struct {
	Error         *string `json:"error,omitempty"`
	MOTD          *string `json:"motd,omitempty"`
	Protocol      *int32  `json:"protocol,omitempty"`
	Version       *string `json:"version,omitempty"`
	PlayersOnline *int32  `json:"players_online,omitempty"`
	PlayersMax    *int32  `json:"players_max,omitempty"`
	ServerID      *string `json:"server_id,omitempty"`
	LevelName     *string `json:"level_name,omitempty"`
	GameModeID    *string `json:"gamemode_id,omitempty"`
	PortV4        *uint16 `json:"port_v4,omitempty"`
	PortV6        *uint16 `json:"port_v6,omitempty"`
}

// GetMinecraftServerStatus 查询 Minecraft 基岩版服务器的状态。
func GetMinecraftServerStatus(host string, port uint16) ServerStatus {
	status := ServerStatus{}

	// 创建 UDP 链接
	conn, err := net.ListenPacket("udp", ":0")
	if err != nil {
		errStr := fmt.Sprintf("无法创建连接: %v", err)
		status.Error = &errStr
		return status
	}
	defer func(conn net.PacketConn) {
		_ = conn.Close()
	}(conn)

	// 设置读取超时
	if err := conn.SetReadDeadline(time.Now().Add(5 * time.Second)); err != nil {
		errStr := fmt.Sprintf("设置读取超时失败: %v", err)
		status.Error = &errStr
		return status
	}

	// 解析目标地址
	targetAddr := fmt.Sprintf("%s:%d", host, port)
	serverAddr, err := net.ResolveUDPAddr("udp", targetAddr)
	if err != nil {
		errStr := fmt.Sprintf("无法解析地址: %v", err)
		status.Error = &errStr
		return status
	}

	// 构造 Ping 数据包
	var buffer bytes.Buffer
	buffer.WriteByte(0x01) // 未连接的 Ping ID

	// 添加时间戳
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	_ = binary.Write(&buffer, binary.BigEndian, timestamp)

	// 添加魔术字节
	buffer.Write(MAGIC)

	// 添加客户端 GUID( 8 字节，共 0 )
	buffer.Write(make([]byte, 8))

	// 发送 Ping 数据包
	if _, err := conn.WriteTo(buffer.Bytes(), serverAddr); err != nil {
		errStr := fmt.Sprintf("发送 Ping 数据包失败: %v", err)
		status.Error = &errStr
		return status
	}

	// 接收响应
	responseBuf := make([]byte, 1500)
	size, _, err := conn.ReadFrom(responseBuf)
	if err != nil {
		errStr := fmt.Sprintf("接收响应失败: %v", err)
		status.Error = &errStr
		return status
	}

	// 验证响应完整性
	if size < 35 || responseBuf[0] != 0x1c {
		errStr := fmt.Sprintf("无效的响应数据包( 大小：%d，ID：%d ) ", size, responseBuf[0])
		status.Error = &errStr
		return status
	}

	// 提取状态字符串
	statusStr := string(responseBuf[35:size])

	// 拆分状态字段
	parts := strings.Split(statusStr, ";")
	if len(parts) < 10 {
		errStr := fmt.Sprintf("状态字段数无效: %d", len(parts))
		status.Error = &errStr
		return status
	}

	// 填充服务器状态
	status.MOTD = &parts[1]

	if protocol, err := strconv.ParseInt(parts[2], 10, 32); err == nil {
		p := int32(protocol)
		status.Protocol = &p
	}

	status.Version = &parts[3]

	if online, err := strconv.ParseInt(parts[4], 10, 32); err == nil {
		o := int32(online)
		status.PlayersOnline = &o
	}

	if maxp, err := strconv.ParseInt(parts[5], 10, 32); err == nil {
		m := int32(maxp)
		status.PlayersMax = &m
	}

	status.ServerID = &parts[6]
	status.LevelName = &parts[7]
	status.GameModeID = &parts[8]

	if len(parts) > 10 {
		if port, err := strconv.ParseUint(parts[10], 10, 16); err == nil {
			p := uint16(port)
			status.PortV4 = &p
		}
	}

	if len(parts) > 11 {
		if port, err := strconv.ParseUint(parts[11], 10, 16); err == nil {
			p := uint16(port)
			status.PortV6 = &p
		}
	}

	return status
}
