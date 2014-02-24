// go的select让使用者可以同时等待多个channel的数据，
// select在连接协程和通道之间是非常强大的工具

package main

import "time"
import "fmt"

func main() {

    // 我们将同时select两个channel
    c1 := make(chan string)
    c2 := make(chan string)

    // Each channel will receive a value after some amount
    // of time, to simulate e.g. blocking RPC operations
    // executing in concurrent goroutines.
    
    // 每一个channel都将在若干秒后收到一个值，使用并发的协程
    // 来模拟阻塞的 RPC 操纵
    go func() {
        time.Sleep(time.Second * 1)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "two"
    }()

    // 我们将使用select来同时等待这两个协程，直到其中一个
    // 有数据返回
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
