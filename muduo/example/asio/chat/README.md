# chat
## 简介
- 简单的聊天服务
- chatserver，聊天服务器
- chatclient，聊天客户端

## 编译
- 使用build脚本可生成chat程序
```shell
$ ./build.sh
```

## 效果
- 首先执行chatserver程序
```shell
$ ./chatserver {端口号}
```
- 然后在两个窗口分别执行chatclient程序，在客户端下面输入文字即可相互通信
```shell
$ ./chatclient {ip} {端口号}
```