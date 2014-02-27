// defer被设计用来确保一个函数调用结束后确定某些行为得到执行
// 通常被用于处理收尾的工作,在别的语言中可能使用`ensure` 和 `finally`
// 语法

package main

import "fmt"
import "os"

// 假设我们需要创建一个文件，然后往里写入数据
// 这时我们可以用`defer`来做些什么
func main() {

    // 我们通过`createFile`得到一个文件对象，然后使用defer
    // 执行`closeFile`来关闭文件这可以确保在写入文件操作完成后
    // 关闭文件的函数得以执行
    f := createFile("/tmp/defer.txt")
    defer closeFile(f)
    writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
    fmt.Println("closing")
    f.Close()
}
