# filetransfer
## 简介
- 文件传输脚本
- pushfileTo，将文件传输到某台机器上
- pullfile，将文件拉取下来

## 安装
- 使用build脚本安装
```shell
$ ./build.sh install
```
- /bin目录下出现如下命令表示安装成功
```shell
$ ls /bin/pu*
/bin/pullfile  /bin/pushfileTo
```
## 用法
- 假如想要将A机器的文件传向B机器，首先在B机器执行如下命令
```shell
$ pullfile
```
- 然后在A机器上执行如下命令
```shell
$ pushfileTo {B机器IP} {要传输的文件或文件夹的名字}
```
## 卸载
- 使用build脚本卸载
```shell
$ ./build.sh uninstall
```
- /bin目录下出现如下命令表示卸载成功
```shell
$ ls /bin/pu*
no matches found: /bin/pu*
```

