# time
## 简介
- timeserver以4个字节的形式发送时间
- timeclient是配套的客户端

## 编译
- 使用build脚本可生成timeserver和timeclient程序
```shell
$ ./build.sh
```

## 效果
- 首先执行timeserver程序
```shell
$ ./timeserver
```
- 然后执行timeclient程序，可以看到如下打印
```shell
$ ./timeclient localhost
20220501 05:35:49.146699Z 4094428 INFO  pid = 4094428 - timeclient.cc:72
20220501 05:35:49.146788Z 4094428 ERROR sockets::fromIpPort - SocketsOps.cc:241
20220501 05:35:49.146802Z 4094428 INFO  TcpClient::TcpClient[TimeClient] - connector 0x9A6960 - TcpClient.cc:69
20220501 05:35:49.146814Z 4094428 INFO  TcpClient::connect[TimeClient] - connecting to 0.0.0.0:2037 - TcpClient.cc:107
20220501 05:35:49.146901Z 4094428 INFO  127.0.0.1:58070 -> 127.0.0.1:2037 is UP - timeclient.cc:41
20220501 05:35:49.147047Z 4094428 INFO  Server time = 1651383349, 20220501 05:35:49.000000 - timeclient.cc:60
20220501 05:35:49.147061Z 4094428 INFO  127.0.0.1:58070 -> 127.0.0.1:2037 is DOWN - timeclient.cc:41
20220501 05:35:49.147093Z 4094428 INFO  TcpClient::~TcpClient[TimeClient] - connector 0x9A6960 - TcpClient.cc:75
```