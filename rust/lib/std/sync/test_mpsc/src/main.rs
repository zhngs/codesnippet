use std::thread;
use std::sync::mpsc::channel;
use std::sync::mpsc::sync_channel;

fn example1() {
    // 创建一个简单的流通道
    let (tx, rx) = channel();
    thread::spawn(move|| {
        tx.send(10).unwrap();
    });
    assert_eq!(rx.recv().unwrap(), 10);
}

fn example2() {
    // 创建一个可以从许多线程一起发送的共享通道，其中 tx 是发送一半 (用于传输的 tx)，rx 是接收一半 (用于接收的 rx)。
    let (tx, rx) = channel();
    for i in 0..10 {
        let tx = tx.clone();
        thread::spawn(move|| {
            tx.send(i).unwrap();
        });
    }
    for _ in 0..10 {
        let j = rx.recv().unwrap();
        assert!(0 <= j && j < 10);
    }
}

fn example3() {
    // 调用 recv() 将返回错误，因为通道已挂起 (或已释放)
    let (tx, rx) = channel::<i32>();
    drop(tx);
    assert!(rx.recv().is_err());
}

fn example4() {
    let (tx, rx) = sync_channel::<i32>(0);
    thread::spawn(move|| {
        // 这将等待父线程开始接收
        tx.send(53).unwrap();
    });
    rx.recv().unwrap();
}

fn example5() {
    let (tx, rx) = sync_channel(3);

    for _ in 0..3 {
        // 这里没有线程和克隆也是一样的，因为仍然会剩下一个 `tx`。
        let tx = tx.clone();
        // 克隆的 tx 丢弃在线程中
        thread::spawn(move || tx.send("ok").unwrap());
    }

    // 删除最后一个发送者停止 `rx` 等待消息。
    // 如果我们将其注释掉，程序将无法完成。
    // 所有需要为 `rx` 排除 `tx` 才能拥有 `Err`。
    drop(tx);

    // 无限接收者等待所有发送者完成。
    while let Ok(msg) = rx.recv() {
        println!("{}", msg);
    }
}

fn main() {
    example1();
    example2();
    example3();
    example4();
    example5();
}
