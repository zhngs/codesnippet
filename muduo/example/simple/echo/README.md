# echo
## 简介
- 回显服务器

## 编译
- 使用build脚本可生成echo程序
```shell
$ ./build.sh
```

## 效果
- 首先执行echo程序
```shell
$ ./echo
```
- 使用netstat查看程序，可以看到程序监听IPv4通配地址的2007端口
```shell
$ netstat -apn | grep echo
tcp        0      0 0.0.0.0:2007            0.0.0.0:*               LISTEN      4096790/./echo      
```
- 使用nc连接该端口，输入hello并按下回车，可以看到hello被回送回来
```shell
$ nc localhost 2009
hello
hello
```