use futures::executor::block_on;

async fn learn_song() {
    println!("learn song");
}

async fn sing_song() {
    println!("sing song");
}

async fn dance() {
    println!("dance");
}

async fn learn_and_sing() {
    // await和block_on的区别在于await不会阻塞当前线程，block_on会
    learn_song().await;
    sing_song().await;
}

async fn async_main() {
    let f1 = learn_and_sing();
    let f2 = dance();

    // join!类似于await，但是可以并行等待多个future
    futures::join!(f1, f2);
}

fn main() {
    block_on(async_main());
}