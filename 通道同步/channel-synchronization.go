// 我们可以使用通道使协程同步执行,下面的例子使用阻塞通道
// 的方式，等待协程执行完毕

package main

import "fmt"
import "time"

// 这个函数是通过协程执行的，done通道用来提醒其它协程，这
// 个函数已经执行完毕
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    // 发送一个值到done通道中表明该函数已经执行完了
    done <- true
}

func main() {

    // 开启一个工作协程，传递一个提示通道
    done := make(chan bool, 1)
    go worker(done)

    // 直到done通道获取数据前，主协程都是阻塞的
    <-done
}
