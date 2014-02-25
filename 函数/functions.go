// 函数是go的核心组成部分，我们看看下面几个简单的例子

package main

import "fmt"

// 下面是一个有两个整形参数的函数，并返回它们的和
// 返回值必须是整形
func plus(a int, b int) int {

	// go要求明确的返回值，不会自动的返回最后一个
	// 表达式
    return a + b
}

func main() {

	// 像你期望的那样调用函数
    // `name(args)`.
    res := plus(1, 2)
    fmt.Println("1+2 =", res)
}

// todo: 合并同类型参数
