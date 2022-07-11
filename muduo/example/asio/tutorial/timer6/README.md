# timer
## 简介
- 多线程回调定时器，但缩小了临界区

## 编译
- 使用build脚本可生成timer程序
```shell
$ ./build.sh
```

## 效果
- 执行timer程序，会打印出1到10
```shell
$ ./timer
Timer 1: 0
Timer 2: 1
Timer 1: 2
Timer 2: 3
Timer 1: 4
Timer 2: 5
Timer 1: 6
Timer 2: 7
Timer 1: 8
Timer 2: 9
Final count is 10
```