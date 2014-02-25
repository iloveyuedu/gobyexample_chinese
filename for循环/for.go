// for是基本的循环语句
// 下面有三种基本的使用方式

package main

import "fmt"

func main() {

    // 最基本的情况，只是用单个条件表达式
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }

    // 经典使用方式 初始化/条件/条件数据更新
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    // 空条件表示这是一个无限循环，只有遇到 `break` 才能
    // 退出循环，或者使用 `return` 结束当前函数调用
    for {
        fmt.Println("loop")
        break
    }
}
