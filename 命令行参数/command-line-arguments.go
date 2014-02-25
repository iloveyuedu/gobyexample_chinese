// [命令行参数](http://en.wikipedia.org/wiki/Command-line_interface#Arguments)
// 是传递参数给程序使用的一种常见的方式
// 比如`go run hello.go` 使用了 `run` 和
// `hello.go`作为go程序的参数

package main

import "os"
import "fmt"

func main() {

    // `os.Args` 提供了直接获取命令参数的入口，注意第一个参数永远是
	// 当前程序的路径,`os.Args[1:]`是所有用户的参数
    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]

	// 你也可以直接通过索引获取参数值
    arg := os.Args[3]

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}
