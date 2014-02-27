// 切片是go里非常重要的数据结果，比数组这类固定
// 长度的序列要强大很多

package main

import "fmt"

func main() {

    // 和数组不太一样，切片只申明类型不用指定长度长度
    // 使用内建的`make`创建一个长度大于0的切片，下面我们
    // 初始化了一个初始长度为3的字符串类型切片,这3个切片元素已经
    // 被初始为0值了
    s := make([]string, 3)
    fmt.Println("emp:", s)

    // 我们可以像数组一样使用索引获取切片元素值
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    // `len` 返回切片的长度
    fmt.Println("len:", len(s))

    // 除了上面那些基础的操作，切片支持比数组更多的操作方式
    // `append`就是另一个例子，它往切片里追加一个新的值，并返回新
    // 的切片,注意我们需要接收append函数的返回值，一个新的切片
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    // 切片也能被`copy`，下面我们创建一个初始长度和切片`s`一样的
    // 另一个切片，并把切片`s`复制到新切片`c`里
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    // 切片支持分片操作，使用`slice[low:high]`语法,比如我们要
    // 获取元素 `s[2]`, `s[3]`, `s[4]`，像下面这样
    l := s[2:5]
    fmt.Println("sl1:", l)

    // 获取`s[5]`元素之前的所有元素
    l = s[:5]
    fmt.Println("sl2:", l)

    // 获取`s[2]`和之后的所有元素
    l = s[2:]
    fmt.Println("sl3:", l)

    // 我们也可以同时申明和初始化一个切片类型变量
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    // 切片也可以是多维的结构，第n维的切片长度也是可变的，而
    // 不像多维数组，长度是固定的
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
