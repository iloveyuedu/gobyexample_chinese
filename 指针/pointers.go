// go支持<em><a href="http://en.wikipedia.org/wiki/Pointer_(computer_programming)">指针</a></em>
// 允许传递值和一片记录的引用

package main

import "fmt"

// 我们将通过两个函数 `zeroval` 和 `zeroptr`来比较指针和普通传递
// `zeroval`有一个`int`类型的参数，所以默认他将会使用值传递，值复制
// 的方式来获取传递的参数
func zeroval(ival int) {
    ival = 0
}

// `zeroptr` 和上面的区别是使用了 `*int` 类型的参数,意味着它是
// 接收一个`int`类型的指针，该函数体内的 `*iptr` 将会对指针进行
// 解引用来指向它真实的地址，这时对它进行赋值操作将会改变该整形
// 指针所指向的整形的真实的值
func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    // `&i` 语法将给出变量`i`在内存的地址，一个指向
    // 变量`i`的指针
    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    // 指针也可以被打印
    fmt.Println("pointer:", &i)
}
