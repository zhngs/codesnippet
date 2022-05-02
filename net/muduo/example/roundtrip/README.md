# roundtrip
## 简介
- 测两台机器间的网络延迟

## 编译
- 使用build脚本可生成roundtrip程序
```shell
$ ./build.sh
```

## 效果
- 执行server端程序
```shell
$ ./roundtrip -s {端口号}
```
- 执行client端程序，打印信息如下
```shell
$ ./roundtrip {ip} {端口号}
20220502 05:56:28.039500Z 126713 INFO  round trip 84 clock error 0 - roundtrip.cc:82
20220502 05:56:28.239528Z 126713 INFO  round trip 70 clock error 6 - roundtrip.cc:82
20220502 05:56:28.439545Z 126713 INFO  round trip 51 clock error 1 - roundtrip.cc:82
20220502 05:56:28.639605Z 126713 INFO  round trip 70 clock error 4 - roundtrip.cc:82
20220502 05:56:28.839674Z 126713 INFO  round trip 104 clock error 5 - roundtrip.cc:82
20220502 05:56:29.039731Z 126713 INFO  round trip 72 clock error 5 - roundtrip.cc:82
20220502 05:56:29.239762Z 126713 INFO  round trip 73 clock error 0 - roundtrip.cc:82
20220502 05:56:29.439824Z 126713 INFO  round trip 97 clock error -10 - roundtrip.cc:82
20220502 05:56:29.639839Z 126713 INFO  round trip 79 clock error 0 - roundtrip.cc:82
```