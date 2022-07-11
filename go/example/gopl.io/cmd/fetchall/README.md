# fetchall

## 简介
- 并发fetch网址

## 编译
- 使用build.sh编译
```shell
$ ./build.sh
```

## 效果
- 执行fetchall程序，根据时间可以发现是并发fetch的
```shell
$ ./fetchall https://www.baidu.com https://www.zhihu.com
0.07s       227    https://www.baidu.com
0.14s     39521    https://www.zhihu.com
0.14s elapsed
```