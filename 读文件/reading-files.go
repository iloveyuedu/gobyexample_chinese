// 读写文件是很常见的任务，我们先看看如何读文件吧

package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

// 读文件大多时候都需要检查错误，使用下面这个辅助函数很有用
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // 或许我们需要一下将文件内容全部读入内存
    dat, err := ioutil.ReadFile("/tmp/dat")
    check(err)
    fmt.Print(string(dat))

    // 如果你想精确的控制读取文件的那块，通过`Open`函数
    // 得到一个 `os.File` 类型的值
    f, err := os.Open("/tmp/dat")

    // 从文件开始读取5个字节的数据，但是得注意实际我们读取了
    // 多少的数据
    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1))

    // 你也可以使用 `Seek` 函数到达文件指定位置，并从那开始读取
    o2, err := f.Seek(6, 0)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

    // `io` 包提供了一些有用的函数来辅助文件的读取，比如
    // 实现上面的至少读取两个字符，使用 `ReadAtLeast` 函数
    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    // go没有提供重置文件指针的函数，但是可以通过
    // `Seek(0, 0)`来实现
    _, err = f.Seek(0, 0)
    check(err)

    // `bufio` 实现了一个可缓冲的读取器，可能会在效率上更好
    // 通过多次读取
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))

    // 关闭文件，通常在打开文件后使用defer马上执行关闭
    f.Close()

}
