// go里写入文件和我们以前看见过的模式很相似

package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // 作为开始，下面展示如何将一个字符串(bytes)写入一个
    // 文件
    d1 := []byte("hello\ngo\n")
    err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
    check(err)

    // 更精细的控制，为写入新建一个文件
    f, err := os.Create("/tmp/dat2")
    check(err)

    // 习惯上，我们在打开一个文件后马上通过defer执行
    // 关闭文件命令
    defer f.Close()

    // 你也可以写入byte切片
    d2 := []byte{115, 111, 109, 101, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("wrote %d bytes\n", n2)

    // 直接写入字符串 `WriteString` 也是可以的
    n3, err := f.WriteString("writes\n")
    fmt.Printf("wrote %d bytes\n", n3)

    // `Sync` 命令会将数据刷到磁盘上
    f.Sync()

    // `bufio` 提供了缓冲写同时我们之前也提到过
    // 缓冲读
    w := bufio.NewWriter(f)
    n4, err := w.WriteString("buffered\n")
    fmt.Printf("wrote %d bytes\n", n4)

    // 使用 `Flush` 命令来确保所有的缓冲操作真实的进行了写操作
    w.Flush()

}
