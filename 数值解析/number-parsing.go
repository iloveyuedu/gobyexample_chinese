// 从字符串中解析出数值是很基础的常见任务，下面
// 我们看看在go里该怎么做

package main

// 内建的 `strconv` 提供了数值解析操作
import "strconv"
import "fmt"

func main() {

    // `ParseFloat`,`64`表示我们在解析时需要保留多少精度
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

    // `ParseInt`, `0` 表示自动推断解析数值的进制，`64`表示结果
    // 使用int64来表示
    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)

    // `ParseInt`能识别16进制的数值字符串
    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)

    // `ParseUint`也是可用的
    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)

    // `Atoi`是一个将字符串转成10进制整形数字的比较方便的一个方法
    k, _ := strconv.Atoi("135")
    fmt.Println(k)

    // 解析函数将对无效的输入返回一个错误
    _, e := strconv.Atoi("wat")
    fmt.Println(e)
}
