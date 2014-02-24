//  是连接多个并发的协程的管道，你可以从一个协程
// 里发送数据到，从另一个协程中读取

package main

import "fmt"

func main() {

    // 通过 `make(chan val-type)` 创建一个新的
    // 是有类型的，指定传递的数据类型
    messages := make(chan string)

    // 发送值到内使用 ` <-` 语法,这里我们发送
    // `"ping"`到`messages`内,这个通道是我们上面
    // 创建的，在一个新的协程里使用
    go func() { messages <- "ping" }()

    // `<-channel`语法 用来接收一个通道内的值，这里我们
    // 接收到`"ping"`消息，我们刚刚在上面发送的消息，并把它打印出来
    msg := <-messages
    fmt.Println(msg)
}
