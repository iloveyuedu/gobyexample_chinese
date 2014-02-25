// 在go里使用`if` 和 `else`分支

package main

import "fmt"

func main() {

    // 基础用法
    if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }

    // 没有else的if
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }

    // 可以包含先决条件，任何在这里申明的变量，在整个循环体内部都是
    // 可以使用的
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}

// 在go里，你不在需要使用圆括号将条件表达式阔起来，单式大括号还是必须的