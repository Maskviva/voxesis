export namespace entity {
	
	export class AppConfig {
	    mc_server_root_path: string;
	    mc_server_proxy_root_path: string;
	    theme: string;
	    qq_bot: boolean;
	    qq_bot_port: number;
	    llonebot_token: string;
	    qq_group: string;
	
	    static createFrom(source: any = {}) {
	        return new AppConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.mc_server_root_path = source["mc_server_root_path"];
	        this.mc_server_proxy_root_path = source["mc_server_proxy_root_path"];
	        this.theme = source["theme"];
	        this.qq_bot = source["qq_bot"];
	        this.qq_bot_port = source["qq_bot_port"];
	        this.llonebot_token = source["llonebot_token"];
	        this.qq_group = source["qq_group"];
	    }
	}
	export class McServerConfig {
	    "server-name": string;
	    gamemode: string;
	    "force-gamemode": string;
	    difficulty: string;
	    "allow-cheats": string;
	    "max-players": string;
	    "online-mode": string;
	    "allow-list": string;
	    "server-port": string;
	    "server-portv6": string;
	    "enable-lan-visibility": string;
	    "view-distance": string;
	    "tick-distance": string;
	    "player-idle-timeout": string;
	    "max-threads": string;
	    "level-name": string;
	    "level-seed": string;
	    "default-player-permission-level": string;
	    "texturepack-required": string;
	    "content-log-file-enabled": string;
	    "compression-threshold": string;
	    "compression-algorithm": string;
	    "server-authoritative-movement-strict": string;
	    "server-authoritative-dismount-strict": string;
	    "server-authoritative-entity-interactions-strict": string;
	    "player-position-acceptance-threshold": string;
	    "player-movement-action-direction-threshold": string;
	    "server-authoritative-block-breaking-pick-range-scalar": string;
	    "chat-restriction": string;
	    "disable-player-interaction": string;
	    "client-side-chunk-generation-enabled": string;
	    "block-network-ids-are-hashes": string;
	    "disable-persona": string;
	    "disable-custom-skins": string;
	    "server-build-radius-ratio": string;
	    "allow-outbound-script-debugging": string;
	    "allow-inbound-script-debugging": string;
	    "script-debugger-auto-attach": string;
	
	    static createFrom(source: any = {}) {
	        return new McServerConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this["server-name"] = source["server-name"];
	        this.gamemode = source["gamemode"];
	        this["force-gamemode"] = source["force-gamemode"];
	        this.difficulty = source["difficulty"];
	        this["allow-cheats"] = source["allow-cheats"];
	        this["max-players"] = source["max-players"];
	        this["online-mode"] = source["online-mode"];
	        this["allow-list"] = source["allow-list"];
	        this["server-port"] = source["server-port"];
	        this["server-portv6"] = source["server-portv6"];
	        this["enable-lan-visibility"] = source["enable-lan-visibility"];
	        this["view-distance"] = source["view-distance"];
	        this["tick-distance"] = source["tick-distance"];
	        this["player-idle-timeout"] = source["player-idle-timeout"];
	        this["max-threads"] = source["max-threads"];
	        this["level-name"] = source["level-name"];
	        this["level-seed"] = source["level-seed"];
	        this["default-player-permission-level"] = source["default-player-permission-level"];
	        this["texturepack-required"] = source["texturepack-required"];
	        this["content-log-file-enabled"] = source["content-log-file-enabled"];
	        this["compression-threshold"] = source["compression-threshold"];
	        this["compression-algorithm"] = source["compression-algorithm"];
	        this["server-authoritative-movement-strict"] = source["server-authoritative-movement-strict"];
	        this["server-authoritative-dismount-strict"] = source["server-authoritative-dismount-strict"];
	        this["server-authoritative-entity-interactions-strict"] = source["server-authoritative-entity-interactions-strict"];
	        this["player-position-acceptance-threshold"] = source["player-position-acceptance-threshold"];
	        this["player-movement-action-direction-threshold"] = source["player-movement-action-direction-threshold"];
	        this["server-authoritative-block-breaking-pick-range-scalar"] = source["server-authoritative-block-breaking-pick-range-scalar"];
	        this["chat-restriction"] = source["chat-restriction"];
	        this["disable-player-interaction"] = source["disable-player-interaction"];
	        this["client-side-chunk-generation-enabled"] = source["client-side-chunk-generation-enabled"];
	        this["block-network-ids-are-hashes"] = source["block-network-ids-are-hashes"];
	        this["disable-persona"] = source["disable-persona"];
	        this["disable-custom-skins"] = source["disable-custom-skins"];
	        this["server-build-radius-ratio"] = source["server-build-radius-ratio"];
	        this["allow-outbound-script-debugging"] = source["allow-outbound-script-debugging"];
	        this["allow-inbound-script-debugging"] = source["allow-inbound-script-debugging"];
	        this["script-debugger-auto-attach"] = source["script-debugger-auto-attach"];
	    }
	}
	export class McServerState {
	    pid: string;
	    cpu: number;
	    memory: number;
	    runTime: string;
	    motd?: string;
	    protocol?: string;
	    version?: string;
	    players_online?: string;
	    players_max?: string;
	    server_id?: string;
	    level_name?: string;
	    gamemode_id?: string;
	    port_v4?: string;
	    port_v6?: string;
	
	    static createFrom(source: any = {}) {
	        return new McServerState(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.pid = source["pid"];
	        this.cpu = source["cpu"];
	        this.memory = source["memory"];
	        this.runTime = source["runTime"];
	        this.motd = source["motd"];
	        this.protocol = source["protocol"];
	        this.version = source["version"];
	        this.players_online = source["players_online"];
	        this.players_max = source["players_max"];
	        this.server_id = source["server_id"];
	        this.level_name = source["level_name"];
	        this.gamemode_id = source["gamemode_id"];
	        this.port_v4 = source["port_v4"];
	        this.port_v6 = source["port_v6"];
	    }
	}
	export class OsState {
	    CpuCores: number;
	    CpuUsage: number;
	    MemoryUsage: number;
	    OsMemory: number;
	
	    static createFrom(source: any = {}) {
	        return new OsState(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.CpuCores = source["CpuCores"];
	        this.CpuUsage = source["CpuUsage"];
	        this.MemoryUsage = source["MemoryUsage"];
	        this.OsMemory = source["OsMemory"];
	    }
	}

}

