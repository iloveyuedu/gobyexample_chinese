// go 支持匿名函数 http://en.wikipedia.org/wiki/Anonymous_function
// http://en.wikipedia.org/wiki/Closure_(computer_science)
// 匿名函数非常有用，当你不想为函数命名时

package main

import "fmt"

// 下面这个函数返回一个匿名函数，并用闭包使用了变量i
func intSeq() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

func main() {

    // 我们调用intSeq，并将结果(函数)付给nextInt变量
    // 由于返回的函数抓住了i这个变量，所以我们每次调用
    // nextInt的时候i的值都会更新
    nextInt := intSeq()

    // 通过多次调用查看nextInt的效果
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    // 重新开启一个计数，但不影响上面的计数
    newInts := intSeq()
    fmt.Println(newInts())
}
