# dup

## 简介
- 找到输入文本中重复的字符串

## 编译
- 使用build.sh编译
```shell
$ ./build.sh
```

## 效果
- 执行dup程序，输入若干行字符串，按下ctrl+d，程序打印出信息
```shell
$ ./dup
hello
hello
world
hello
^D
3       hello
```
- 或者读取文件中的内容
```shell
$ ./dup test.txt
2       world
3       hello
```