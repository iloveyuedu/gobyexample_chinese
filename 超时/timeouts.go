// 超时设置 在连接外部资源的时候是显得非常重要的，通常都会限制
// 执行时间，用go实现超时是非常简单的，得益于go的channel和select

package main

import "time"
import "fmt"

func main() {

    // 在我们的例子里，假设我们需要连接一个外部资源，两秒之后
    // 把结果存储到channel c1内。
    c1 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c1 <- "result 1"
    }()

    // 下面通过 `select` 实现一个超时
    // `res := <-c1` 等待结果的返回，`<-Time.After`超时设定
    // 如果超过1s c1没有返回，那么将执行超时case
    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(time.Second * 1):
        fmt.Println("timeout 1")
    }

    // 如果我们允许更长的超时，比如3s，那么我们会收到c2成功
    // 的返回，我们将打印该结果
    c2 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(time.Second * 3):
        fmt.Println("timeout 2")
    }
}

// todo: cancellation?
