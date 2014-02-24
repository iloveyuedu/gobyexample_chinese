// go 支持闭包
// <a href="http://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>recursive functions</em></a>.
// 下面是经典的阶乘

package main

import "fmt"

// 这个阶乘函数调用了它自身直到到达fact(0)
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func main() {
    fmt.Println(fact(7))
}
