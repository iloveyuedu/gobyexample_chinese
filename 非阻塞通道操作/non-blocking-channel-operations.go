// 使用`select`实现通道非阻塞的发送，接收，甚至多通道
// 同时非阻塞

package main

import "fmt"

func main() {
    messages := make(chan string)
    signals := make(chan bool)

    // 下面是一个非阻塞的接收，如果在`messages`通道中已经有一个
    // 可用的值，那么将会执行第一个case，并获取那个值，否者会
    // 立即执行default语句
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

    // 非阻塞的发送和上面的接收也差不多
    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

    // 我们在执行 `default` 前可以用多个case实现多路通道的非阻塞
    // 下面是非阻塞的从 `messages` and `signals` 接收值
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}
