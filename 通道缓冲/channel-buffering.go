// 默认 通道 是不能缓冲的，意味着向通道内发送数据后，必须有
// 正确的接收，否者就会阻塞，缓冲通道接收指定数量的数据发送，即使
// 通道数据没有被取出

package main

import "fmt"

func main() {

	// 这里 我们申明了一个可缓冲2个字符串的通道
    messages := make(chan string, 2)

    // Because this channel is buffered, we can send these
    // values into the channel without a corresponding
    // concurrent receive.
	// 因为这个通道是可缓冲的，我们可以发送多个值到通道内
	// 即使没有正确的接收响应
    messages <- "buffered"
    messages <- "channel"

	// 我们可以正常的接收这两个值
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
