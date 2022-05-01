# discard
## 简介
- 丢弃所有收到的数据

## 编译
- 使用build脚本可生成discard程序
```shell
$ ./build.sh
```

## 效果
- 首先执行discard程序
```shell
$ ./discard
20220501 04:59:28.198760Z 4086414 INFO  pid = 4086414 - main.cc:13
```
- 使用netstat查看程序，可以看到程序监听IPv4通配地址的2009端口
```shell
$ netstat -apn | grep discard
tcp        0      0 0.0.0.0:2009            0.0.0.0:*               LISTEN      4086414/./discard
```
- 使用nc连接该端口，输入hello并按下回车，并按下ctrl+c关闭nc
```shell
$ nc localhost 2009
hello
^C
```
- 回到discard程序下面，可以看到如下打印信息
```shell
$ ./discard
20220501 04:59:28.198760Z 4086414 INFO  pid = 4086414 - main.cc:13
20220501 05:03:17.238066Z 4086414 INFO  TcpServer::newConnection [DiscardServer] - new connection [DiscardServer-0.0.0.0:2009#1] from 127.0.0.1:45654 - TcpServer.cc:80
20220501 05:03:17.238106Z 4086414 INFO  DiscardServer - 127.0.0.1:45654 -> 127.0.0.1:2009 is UP - discard.cc:25
20220501 05:03:21.180770Z 4086414 INFO  DiscardServer-0.0.0.0:2009#1 discards 6 bytes received at 1651381401.180744 - discard.cc:35
20220501 05:04:54.277390Z 4086414 INFO  DiscardServer - 127.0.0.1:45654 -> 127.0.0.1:2009 is DOWN - discard.cc:25
20220501 05:04:54.277413Z 4086414 INFO  TcpServer::removeConnectionInLoop [DiscardServer] - connection DiscardServer-0.0.0.0:2009#1 - TcpServer.cc:109
```
- 可以看到打印信息中"discards 6 bytes"，正好是"hello\n"字符串的长度