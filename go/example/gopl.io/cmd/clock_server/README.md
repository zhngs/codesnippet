# clock server

## 简介
- 并发时钟服务器

## 编译
- 使用build.sh进行编译
```shell
$ ./build.sh
```

## 效果
- 首先执行clock_server程序
```shell
# 窗口1
$ ./clock_server
```
- 然后通过nc连接clock_server
```shell
# 窗口2
$ nc localhost 7878
10:34:54
10:34:55
10:34:56
10:34:57
```