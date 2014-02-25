// 协程是轻量级的线程

package main

import "fmt"

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

    // 假设我们有一个函数调用 `f(s)`,下面我们正常的同步的
    // 方式调用它
    f("direct")

    // 通过协程来调用这个函数使用`go f(s)`语法,新的协程
    // 将会并发的执行这个函数调用
    go f("goroutine")

    // You can also start a goroutine for an anonymous
    // function call.
    // 也可以将匿名函数调用放到一个新协程里
    go func(msg string) {
        fmt.Println(msg)
    }("going")

    // 我们的两个新开的协程是以异步的方式在执行，所以我们需要等
    // 他们执行完成，`Scanln`要求我们按一个键，来延迟程序退出
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}
