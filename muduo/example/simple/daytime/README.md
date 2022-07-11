# daytime
## 简介
- 向客户端发送字符串形式的时间

## 编译
- 使用build脚本可生成daytime程序
```shell
$ ./build.sh
```

## 效果
- 首先执行daytime程序
```shell
$ ./daytime
```
- 使用netstat查看程序，可以看到程序监听IPv4通配地址的2013端口
```shell
$ netstat -apn | grep daytime
tcp        0      0 0.0.0.0:2013            0.0.0.0:*               LISTEN      4091075/./daytime
```
- 使用nc连接该端口，可以看到打印出时间
```shell
$ nc localhost 2013
20220501 05:21:26.854018
```