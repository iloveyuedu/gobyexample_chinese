// 我们如果想要在未来的某个时间执行go代码，或者是间隔执行某些
// 任务，go内建了定时器和间隔执行 使这些要求变得简单，我们首先
// 看如何使用定时器，然后再看间隔执行

package main

import "time"
import "fmt"

func main() {

    // 倒计时代表了未来的单个时间，你可以指定需要等待多久执行
    // 它提供了一个通道在未来指定的时间，下面的这个定时器将会
    // 等待两秒执行
    timer1 := time.NewTimer(time.Second * 2)

    // `<-timer1.C`将阻塞知道定时器的通道 知道定时器通道`C`发送出一个值表明
    // 该定时器已经到达指定的时间
    <-timer1.C
    fmt.Println("Timer 1 expired")

    // 如果你只是想等待一段时间，你可以使用`time.Sleep`
    // 而使用定时器的一个原因，你可以在过期之前取消定时
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
}
