// go内建支持函数多返回值，在go里这是非常地道的东西，比如
// 返回两个结果，一个是函数结果，一个用来表示该函数执行
// 是否有异常

package main

import "fmt"

// `(int, int)`在函数结构里表示该函数返回两个`int`s类型的值
func vals() (int, int) {
    return 3, 7
}

func main() {

    // 下面我们使用同时多赋值操作接收返回的两个值
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    // 如果你想丢弃某个返回值，使用`_`来占位即可
    _, c := vals()
    fmt.Println(c)
}

// todo: 命名返回值
// todo: naked returns
