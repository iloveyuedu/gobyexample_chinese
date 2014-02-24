// 定时器用来在未来执行一次，而ticker可以在指定时间间隔
// 重复执行某件事,下面就是如何使用ticker间隔有规律的重复执行的示例

package main

import "time"
import "fmt"

func main() {

    // tickers和定时器使用相似机制：一个通道发送值，我们
    // 使用内建的range语法迭代该通道，该通道每隔500ms会产生
    // 一个新的值
    ticker := time.NewTicker(time.Millisecond * 500)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at", t)
        }
    }()

    // tickers可以像定时器那样手动停止,一旦ticker停止了
    // ticker的C通道不会再发送出任何值,我们在1500ms后停止tickers
    time.Sleep(time.Millisecond * 1500)
    ticker.Stop()
    fmt.Println("Ticker stopped")
}
