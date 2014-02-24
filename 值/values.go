// go有多种类型包括 字符串 整形 浮点 布尔 等等
// 下面是几种简单的示例

package main

import "fmt"

func main() {

    // 字符串 可以通过+来进行连接
    fmt.Println("go" + "lang")

    // 整形和浮点
    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    // 布尔 布尔操作符
    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)
}
