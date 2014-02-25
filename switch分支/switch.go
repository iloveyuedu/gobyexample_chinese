// switch 条件分支

package main

import "fmt"
import "time"

func main() {

    // 下面是简单的 `switch` 用法.
    i := 2
    fmt.Print("write ", i, " as ")
    switch i {
    case 1:
        fmt.Println("one")
    case 2:
        fmt.Println("two")
    case 3:
        fmt.Println("three")
    }

	// 可以使用逗号来分割多个可能的表达式，同时我们使用
	// 了default语句
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("it's the weekend")
    default:
        fmt.Println("it's a weekday")
    }

	// `switch`不带参数可以用来实现if/else的效果
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("it's before noon")
    default:
        fmt.Println("it's after noon")
    }
}

// todo: type switches
