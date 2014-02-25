// [环境变量](http://en.wikipedia.org/wiki/Environment_variable)
// 是unix下一种全局配置的机制 [conveying configuration
// information to Unix programs](http://www.12factor.net/config).
// 下面我们看看如何设置，获取，列出所有的环境变量

package main

import "os"
import "strings"
import "fmt"

func main() {

    // 通过`os.Setenv`设置一个键值对，使用`os.Getenv`获取环境变量
    // 如果不存在，将返回一个空的字符串
    os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR"))

    // 使用 `os.Environ`来获取所有的环境键值对变量,它返回
    // 一个字符串切片，你可以使用 `strings.Split`将键和值分开
    fmt.Println()
    for _, e := range os.Environ() {
        pair := strings.Split(e, "=")
        fmt.Println(pair[0])
    }
}
