# chargen
## 简介
- chargen服务端不断发送字符
- chargenclient为配套客户端

## 编译
- 使用build脚本可生成discard程序
```shell
$ ./build.sh
```

## 效果
- 首先执行chargen程序
```shell
$ ./chargen
```
- 然后执行chargenclient程序
```shell
$ ./chargenclient localhost
```
- 回到chargen程序下面，可以看到如下打印信息
```shell
$ ./chargen
20220501 06:00:58.540801Z 4100152 INFO  pid = 4100152 - main.cc:13
0.000 MiB/s
0.000 MiB/s
0.000 MiB/s
0.000 MiB/s
20220501 06:01:12.146958Z 4100152 INFO  TcpServer::newConnection [ChargenServer] - new connection [ChargenServer-0.0.0.0:2078#1] from 127.0.0.1:55030 - TcpServer.cc:80
20220501 06:01:12.146995Z 4100152 INFO  ChargenServer - 127.0.0.1:55030 -> 127.0.0.1:2078 is UP - chargen.cc:49
351.172 MiB/s
744.300 MiB/s
748.032 MiB/s
747.426 MiB/s
743.460 MiB/s
```