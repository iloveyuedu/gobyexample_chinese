// go 支持 字符，字符串，布尔，数字 常量

package main

import "fmt"
import "math"

// `const`定义常量
const s string = "constant"

func main() {
    fmt.Println(s)

    // A `const` statement can appear anywhere a `var`
    // statement can.
    // 一个常量表达式可以出现在任何变量可以出现的地方
    const n = 500000000

    // 常量表达式可以指定精度，也可以进行一些计算
    const d = 3e20 / n
    fmt.Println(d)

    // 一个数值的常量是没有类型的，你需要明确的对它进行转换
    fmt.Println(int64(d))

    // 一个数值常量在不同的上下文会被转成需要的类型
    // 比如一个函数，`math.Sin`需要一个`float64`类型的参数
    fmt.Println(math.Sin(n))
}
