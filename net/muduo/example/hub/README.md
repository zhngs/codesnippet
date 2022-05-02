# hub
## 简介
- 简单的应用层广播
- hub - a server for broadcasting
- pubsub - a client library of hub
- pub - a command line tool for publishing content on a topic
- sub - a demo tool for subscribing a topic

## 编译
- 使用build脚本可生成hub程序
```shell
$ ./build.sh
```

## 效果
- 首先执行hub程序
```shell
# 窗口1
$ ./hub 9999
```
- 然后订阅一个topic
```shell
# 窗口2
$ ./sub localhost:9999 hello
```
- 最后发布消息，在窗口2就可以看到订阅的消息
```shell
# 窗口3
$ ./pub localhost:9999 hello "hello, world!"
```