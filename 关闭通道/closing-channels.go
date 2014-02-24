// 关闭通道意味着不会再有值会被发送到通道离了，或许这对于
// 通道接收者很有用
package main

import "fmt"

// 这里，我们将使用一个 `jobs` 通道来纷发工作,工作协程
// 会从jobs通道中读取任务并执行，当我们没有工作需要执行时，我们
// 将关闭该通道。
func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    // 下面是工作协程，它重复的从jobs通道中接收数据并放入j变量中
    // 当通道jobs关闭后并且jobs通道读取完成后more变量将会为false,
    // 我们使用done通道来表示所有的工作都已经执行完毕
    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    // 发送3个工作任务到jobs通道，然后关闭它
    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    // 我们使用同步channel来避免程序过早退出
    <-done
}
