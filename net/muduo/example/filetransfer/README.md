# filetransfer
## 简介
- 并发文件服务器

## 编译
- 使用build脚本可生成uploadfile程序
```shell
$ ./build.sh
```

## 效果
- 首先执行uploadfile程序
```shell
$ ./uploadfile README.md
```
- 使用netstat查看程序，可以看到程序监听IPv4通配地址的2021端口
```shell
$ netstat -apn | grep upload
tcp        0      0 0.0.0.0:2021            0.0.0.0:*               LISTEN      4108366/./uploadfil
```
- 使用nc连接该端口
```shell
$ nc localhost 2021 > README.md
```
- 回到uploadfile程序下面，可以看到如下打印信息
```shell
$ ./uploadfile
20220501 06:37:13.455806Z 4108366 INFO  pid = 4108366 - download3.cc:66
20220501 06:41:14.382491Z 4108366 INFO  TcpServer::newConnection [FileServer] - new connection [FileServer-0.0.0.0:2021#1] from 127.0.0.1:33186 - TcpServer.cc:80
20220501 06:41:14.382535Z 4108366 INFO  FileServer - 127.0.0.1:33186 -> 127.0.0.1:2021 is UP - download3.cc:22
20220501 06:41:14.382549Z 4108366 INFO  FileServer - Sending file README.md to 127.0.0.1:33186 - download3.cc:27
20220501 06:41:14.382618Z 4108366 INFO  FileServer - done - download3.cc:60
```
- 当看到打印信息中的"FileServer - done"说明文件传输完成
- 然后回到nc窗口按下ctrl+c即可看到传输的文件