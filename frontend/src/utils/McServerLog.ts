import { ref } from 'vue'

interface Log {
    time: string | null;
    level: string;
    message: string;
    fullMessage?: string;
    raw: string;
}

// // 日志解析逻辑 (保持不变)
// const createLogParser = () => {
//     let startTime: number | null = null;
//     let lastKnownTime: string | null = null;
//     const parseSingleLine = (logLine: string): Log | null => {
//         const standardRegex = /^\[(\d{4}-\d{2}-\d{2}) (\d{2}):(\d{2}):(\d{2}):(\d{3}) (\w+)](.*)$/;
//         let match = logLine.match(standardRegex);
//         if (!match) {
//             return { time: lastKnownTime || "??:??.???", level: 'UNKNOWN', message: logLine.trim(), raw: logLine };
//         }
//         const [, date, hours, minutes, seconds, ms, level, content] = match;
//         const timestamp = `${date}T${hours}:${minutes}:${seconds}.${ms.padStart(3, '0')}Z`;
//         const currentTime = new Date(timestamp).getTime();
//         if (!startTime) startTime = currentTime;
//         const timeDiff = currentTime - startTime;
//         const mins = Math.floor(timeDiff / 60000);
//         const secs = Math.floor((timeDiff % 60000) / 1000);
//         const millis = timeDiff % 1000;
//         const formattedTime = `${mins.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}.${millis.toString().padStart(3, '0')}`;
//         lastKnownTime = formattedTime;
//         let message = content.trim();
//         const MAX_LENGTH = 120;
//         const displayMessage = message.length > MAX_LENGTH ? `${message.substring(0, MAX_LENGTH)}...` : message;
//         return { time: formattedTime, level, message: displayMessage, fullMessage: message, raw: logLine };
//     };
//     return (logLine: any): Log[] => {
//         const lines = logLine.split(/\r?\n/).filter((line: string) => line.trim());
//         return lines.map((line: any) => parseSingleLine(line)).filter((log: Log | null): log is Log => log !== null);
//     };
// };
//
// const parseLogLine = createLogParser();
//
// export const logs = ref<Log[]>([]);
//
// export let unlistenFn: UnlistenFn | null = null;
//
// // 创建监听器但不立即取消它
// listen('mc_server_output', (event) => {
//     console.log('收到服务器输出:', event.payload);
//     const parsedLogs = parseLogLine(event.payload);
//     if (parsedLogs.length > 0) {
//         logs.value.push(...parsedLogs);
//     }
// }).then(unlisten => {
//     // 保存取消监听的函数
//     unlistenFn = unlisten;
// });

// 提供一个清理函数
export const cleanup = () => {
    // if (unlistenFn) {
    //     unlistenFn();
    //     unlistenFn = null;
    // }
};
