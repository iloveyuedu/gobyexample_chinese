// go为字符串格式化提供了十分优秀的控制
// 除了`printf`，这里还有大量的与字符串格式化相关的操作

package main

import "fmt"
import "os"

type point struct {
    x, y int
}

func main() {

    // go为变量打印提供了一些好用的方法，比如，下面
    // 就是打印一个我们自定义的`point`类型的结构
    p := point{1, 2}
    fmt.Printf("%v\n", p)

    // 如果变量是一个结构体，`%+v` 将把结构体的字段也打印出来
    fmt.Printf("%+v\n", p)

    // `%#v`将会打印出该值的语法表达式，产生这个值的源代码片段
    fmt.Printf("%#v\n", p)

    // 通过`%T`打印出一个值的类型
    fmt.Printf("%T\n", p)

    // 格式化布尔变量
    fmt.Printf("%t\n", true)

    // 这里还有大量的选项用来格式化整形，`%d`用来
    // 格式化10进制
    fmt.Printf("%d\n", 123)

    // 下面打印变量的二进制表示
    fmt.Printf("%b\n", 14)

    // 打印出一个整形所代表的字符
    fmt.Printf("%c\n", 33)

    // `%x`用来进行16进制编码
    fmt.Printf("%x\n", 456)

    // 这里也有多个格式化浮点数的方式，格式化成10进制使用`%f`
    fmt.Printf("%f\n", 78.9)

    // `%e` 和 `%E`用来将浮点数格式化成科学计数法的格式
    fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)

    // `%s`用来格式化最基本的字符串
    fmt.Printf("%s\n", "\"string\"")

    // 用双引号包围字符串使用`%q`选项
    fmt.Printf("%q\n", "\"string\"")

    // 我们上面对整形使用了`%x`，也可以将字符串渲染成16进制
    // 每一字节将会输出两个字符串
    fmt.Printf("%x\n", "hex this")

    // 打印一个指针的表达式使用`%p`选项
    fmt.Printf("%p\n", &p)

    // 在格式化数字的时候通常你都会去控制它的长度，在%后指定
    // 长度，默认打印出的数字将会填补空白达到指定的长度,左对齐
    fmt.Printf("|%6d|%6d|\n", 12, 345)

    // 你也可以指定浮点数的打印长度，通常会限制精确度
    // 那么你可以使用下面的语法
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

    // 反向填充使用`-`标记,右对其
    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

    // 你也可以为字符串指定输出长度，在表格对其时可能非常有用
    // 下面是右对齐
    fmt.Printf("|%6s|%6s|\n", "foo", "b")

    // 使用 `-` 标记实现左对齐
    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

    // 目前我们已经见过`Printf`,他将结果打印到 `os.Stdout`里
    // `Sprintf` 将会返回格式化之后的字符串
    s := fmt.Sprintf("a %s", "string")
    fmt.Println(s)

    // 你也可以使用`Fprintf`将结果打印到`io.Writers`而不是一直打印到`os.Stdout`
    fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
