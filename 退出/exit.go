// 使用 `os.Exit` 实现马上退出程序，并同时返回一个状态。

package main

import "fmt"
import "os"

func main() {

	// 使用 `os.Exit`退出`defer`s 语句不会执行
	// 所以 `fmt.Println` 永远不会执行
    defer fmt.Println("!")

	// 退出并返回状态码3
    os.Exit(3)
}

// 注意 go不像c，go不会从像c那样返回(return)一个整形,如果你想
// 以一个非零的状态(non-zero status)退出,你应该使用 `os.Exit`。