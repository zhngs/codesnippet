# filetransfer
## 简介
- 文件传输脚本
- pushfileTo，将文件传输到某台机器上
- pullfile，将文件拉取下来

## 用法
- 假如想要将A机器的文件传向B机器，首先在B机器执行如下命令
```shell
$ ./pullfile
```
- 然后在A机器上执行如下命令
```shell
$ ./pushfileTo {B机器IP} {要传输的文件或文件夹的名字}
```