# mpsc

## 简介
- 多生产者，单消费者 FIFO 队列通信原语

## 编译
- 使用build.sh编译
```shell
$ ./build.sh
```

## 要点
该模块通过通道提供基于消息的通信，具体定义为以下三种类型：
- Sender
- SyncSender
- Receiver

Sender 或 SyncSender 用于将数据发送到 Receiver。 两个发送者都是可克隆的 (multi-producer)，因此许多线程可以同时发送到一个接收者 (single-consumer)

这些通道有两种区别：

- 异步，无限缓冲的通道。 channel 函数将返回 (Sender, Receiver) 元组，其中所有发送都是异步的 (它们从不阻塞)。 通道在概念上具有无限的缓冲区。

- 同步的有界通道。 sync_channel 函数将返回 (SyncSender, Receiver) 元组，其中待处理消息的存储是固定大小的预分配缓冲区。 所有的发送都将被阻塞，直到有可用的缓冲区空间。 请注意，允许的界限为 0，从而使通道成为 “rendezvous” 通道，每个发送者在该通道上原子地将消息传递给接收者。