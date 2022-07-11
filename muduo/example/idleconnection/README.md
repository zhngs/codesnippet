# idleconnection
## 简介
- 对连接进行超时管理

## 编译
- 使用build脚本可生成echo_idleconn程序
```shell
$ ./build.sh
```

## 效果
- 首先执行echo_idleconn程序，指定超时时间为8s
```shell
# 窗口1
$ ./echo_idleconn 8
```
- 然后使用nc连接echo_idleconn
```shell
# 窗口2
$ nc localhost 2007
```
- 窗口1就会打印出时间轮盘
```shell
[0] len = 0 : 
[1] len = 0 : 
[2] len = 0 : 
[3] len = 0 : 
[4] len = 0 : 
[5] len = 0 : 
[6] len = 0 : 
[7] len = 1 : 0xac92d0(2), 
20220502 06:59:59.685617Z 139838 INFO  size = 8 - echo.cc:82
[0] len = 0 : 
[1] len = 0 : 
[2] len = 0 : 
[3] len = 0 : 
[4] len = 0 : 
[5] len = 0 : 
[6] len = 1 : 0xac92d0(1), 
[7] len = 0 : 
```