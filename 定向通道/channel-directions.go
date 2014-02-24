// 当我们使用通道作为函数参数时，可以指定通道只能用于发送
// 或接收值，这种指定通道方向的做法可以让程序更加健壮

package main

import "fmt"

// 这个 `ping` 函数限定该通道只能用于接收值
func ping(pings chan<- string, msg string) {
    pings <- msg
}

// 这个 `pong` 函数限定pings通道只能用于发送值,pongs通道
// 只能用于接收值
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
