// [变参函数](http://en.wikipedia.org/wiki/Variadic_function)
// 能够通过任意个数的参数进行调用,该函数的最后一个参数将容纳所有
// “多余”的变量, 比如`fmt.Println`就是一个常见的变参函数
package main

import "fmt"

// 下面的函数将接受任意数量的`ints`类型的参数
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {

    // 变参函数和正常函数调用没啥区别，只不可可以有任意多个参数
    sum(1, 2)
    sum(1, 2, 3)

    // 如果已经有一个参数切片了，将它们应用到变参函数使用
    // `func(slice...)`语法，就像下面这样
    nums := []int{1, 2, 3, 4}
    sum(nums...)
}
