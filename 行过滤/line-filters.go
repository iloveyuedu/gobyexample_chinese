// 行过滤是一种普通的程序，它从stdin读入数据，然后输出处理之后的数据到stdout
// `grep` 和 `sed`就是这种行过滤程序

// 下面的使用go实现了一个行过滤程序，它把输入的文本转为大写
// 你也可以使用这种模式实现自己的行过滤程序
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {

    // 用缓冲的scanner去包裹非缓冲的 `os.Stdin` 将提供给我们一个方便
    // 的 `Scan` 方法,用它来获取一行的输入
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        // `Text`函数返回当前标记，这里是从输入里获取下一行
        ucl := strings.ToUpper(scanner.Text())

        // 输出大写的输入
        fmt.Println(ucl)
    }

    // 检查在`Scan`操作中是否有错误发生, 在`Scan`时即使有错误
    // 也不会报告，所以需要在这里处理
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
