package entity

type ProcessState struct {
	Pid     string  `json:"pid"`
	Cpu     float64 `json:"cpu"`
	Memory  float64 `json:"memory"`
	RunTime string  `json:"runTime"`
}

type McServerState struct {
	ProcessState
	MOTD          string `json:"motd,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
	Version       string `json:"version,omitempty"`
	PlayersOnline string `json:"players_online,omitempty"`
	PlayersMax    string `json:"players_max,omitempty"`
	ServerID      string `json:"server_id,omitempty"`
	LevelName     string `json:"level_name,omitempty"`
	GameModeID    string `json:"gamemode_id,omitempty"`
	PortV4        string `json:"port_v4,omitempty"`
	PortV6        string `json:"port_v6,omitempty"`
}
