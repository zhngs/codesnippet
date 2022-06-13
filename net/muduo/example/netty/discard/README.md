# discard
## 简介
- discard_server丢弃所有收到的数据，带流量统计
- discard_client为配套的客户端

## 编译
- 使用build脚本可生成discard程序
```shell
$ ./build.sh
```

## 效果
- 首先执行discard_server程序
```shell
# 窗口1
$ ./discard_server
```
- 然后执行discard_client
```shell
# 窗口2
$ ./discard_client localhost 256
```
- 回到窗口1，可以看到如下打印
```shell
2.210 MiB/s 2.075 Ki Msgs/s 1090.61 bytes per msg
4.166 MiB/s 4.209 Ki Msgs/s 1013.65 bytes per msg
4.989 MiB/s 4.871 Ki Msgs/s 1048.79 bytes per msg
4.735 MiB/s 4.545 Ki Msgs/s 1066.94 bytes per msg
4.597 MiB/s 4.585 Ki Msgs/s 1026.62 bytes per msg
4.331 MiB/s 3.326 Ki Msgs/s 1333.57 bytes per msg
4.131 MiB/s 3.769 Ki Msgs/s 1122.32 bytes per msg
4.184 MiB/s 4.497 Ki Msgs/s 952.69 bytes per msg
4.838 MiB/s 4.075 Ki Msgs/s 1215.87 bytes per msg
4.789 MiB/s 5.076 Ki Msgs/s 966.05 bytes per msg
5.107 MiB/s 6.281 Ki Msgs/s 832.59 bytes per msg
4.551 MiB/s 4.080 Ki Msgs/s 1142.26 bytes per msg
4.466 MiB/s 3.569 Ki Msgs/s 1281.33 bytes per msg
```

## 测试
- 本机单线程情况下，测试不同消息大小对带宽的影响
```shell
message size       bandwidth
128                 13 Mib/s
256                 26 Mib/s
512                 53 Mib/s
1024                97 Mib/s
2048                209 Mib/s
4096                413 Mib/s
65536               3814 Mib/s
```
- 可以发现，带宽和消息大小成正比，并且经过验证，cpu的主要开销在内核态
- 不难得出结论：单位时间内发出的消息数是一定的，原因是单位时间内系统调用的次数一定，且消息的大小对系统调用的影响不大