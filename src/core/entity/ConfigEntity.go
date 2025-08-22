package entity

type AppConfig struct {
	McServerRootPath      string `json:"mc_server_root_path"`
	McServerProxyRootPath string `json:"mc_server_proxy_root_path"`
	Theme                 string `json:"theme"`
	QQBot                 bool   `json:"qq_bot"`
	QQBotPort             int64  `json:"qq_bot_port"`
	LlonebotToken         string `json:"llonebot_token"`
	QQGroup               string `json:"qq_group"`
}

type McServerConfig struct {
	ServerName                                      string `json:"server-name"`
	Gamemode                                        string `json:"gamemode"`
	ForceGamemode                                   string `json:"force-gamemode"`
	Difficulty                                      string `json:"difficulty"`
	AllowCheats                                     string `json:"allow-cheats"`
	MaxPlayers                                      string `json:"max-players"`
	OnlineMode                                      string `json:"online-mode"`
	AllowList                                       string `json:"allow-list"`
	ServerPort                                      string `json:"server-port"`
	ServerPortv6                                    string `json:"server-portv6"`
	EnableLanVisibility                             string `json:"enable-lan-visibility"`
	ViewDistance                                    string `json:"view-distance"`
	TickDistance                                    string `json:"tick-distance"`
	PlayerIdleTimeout                               string `json:"player-idle-timeout"`
	MaxThreads                                      string `json:"max-threads"`
	LevelName                                       string `json:"level-name"`
	LevelSeed                                       string `json:"level-seed"`
	DefaultPlayerPermissionLevel                    string `json:"default-player-permission-level"`
	TexturepackRequired                             string `json:"texturepack-required"`
	ContentLogFileEnabled                           string `json:"content-log-file-enabled"`
	CompressionThreshold                            string `json:"compression-threshold"`
	CompressionAlgorithm                            string `json:"compression-algorithm"`
	ServerAuthoritativeMovementStrict               string `json:"server-authoritative-movement-strict"`
	ServerAuthoritativeDismountStrict               string `json:"server-authoritative-dismount-strict"`
	ServerAuthoritativeEntityInteractionsStrict     string `json:"server-authoritative-entity-interactions-strict"`
	PlayerPositionAcceptanceThreshold               string `json:"player-position-acceptance-threshold"`
	PlayerMovementActionDirectionThreshold          string `json:"player-movement-action-direction-threshold"`
	ServerAuthoritativeBlockBreakingPickRangeScalar string `json:"server-authoritative-block-breaking-pick-range-scalar"`
	ChatRestriction                                 string `json:"chat-restriction"`
	DisablePlayerInteraction                        string `json:"disable-player-interaction"`
	ClientSideChunkGenerationEnabled                string `json:"client-side-chunk-generation-enabled"`
	BlockNetworkIdsAreHashes                        string `json:"block-network-ids-are-hashes"`
	DisablePersona                                  string `json:"disable-persona"`
	DisableCustomSkins                              string `json:"disable-custom-skins"`
	ServerBuildRadiusRatio                          string `json:"server-build-radius-ratio"`
	AllowOutboundScriptDebugging                    string `json:"allow-outbound-script-debugging"`
	AllowInboundScriptDebugging                     string `json:"allow-inbound-script-debugging"`
	ScriptDebuggerAutoAttach                        string `json:"script-debugger-auto-attach"`
}
