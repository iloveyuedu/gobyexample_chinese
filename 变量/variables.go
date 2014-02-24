// 在go里，编译器会检查变量的类型和函数调用

package main

import "fmt"

func main() {

    // `var` 关键字定义1个或多个变量
    var a string = "initial"
    fmt.Println(a)

    // 同时定义多个变量.
    var b, c int = 1, 2
    fmt.Println(b, c)

    // Go会通过初始值推断出变量的类型
    var d = true
    fmt.Println(d)

	// 为初始化的变量将会默认初始化为该类型的0值，比如
	// `int`类型的0值是`0`。
    var e int
    fmt.Println(e)

	// `:=` 语法是定义并初始化一个变量的快捷方式
	// 下面的示例效果和`var f string = "short"`一样
    f := "short"
    fmt.Println(f)
}
