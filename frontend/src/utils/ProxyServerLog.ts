import {ref} from 'vue'

export const logs = ref<string[]>([]);

// export let unlistenFn: UnlistenFn | null = null;

// 创建监听器但不立即取消它
// listen('proxy_server_output', (event) => {
//     console.log('收到代理服务器输出:', event.payload);
//     logs.value.push(event.payload as string);
// }).then(unlisten => {
//     // 保存取消监听的函数
//     unlistenFn = unlisten;
// });

// 提供一个清理函数
export const cleanup = () => {
    // if (unlistenFn) {s
    //     unlistenFn();s
    //     unlistenFn = snull;
    // }s
};