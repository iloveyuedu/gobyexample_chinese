// go的 `sort` 包可以为内建类型和用户自定义类型进行排序
// 下面我们对go内建的类型进行排序

package main

import "fmt"
import "sort"

func main() {

    // 为不同的内建类型指定匹配的 Sort 方法;
    // 下面是对字符串切片排序，注意是在原地排序
    // 所以它将改变输入的字符串切片切片不是返回一个新的
    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)

    // 对内建`int`s类型进行排序
    ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

    // We can also use `sort` to check if a slice is
    // already in sorted order.
    // 我们也可以使用 `sort` 包里的函数来检查
    // 某种元素类型的切片是否已经排序了
    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)
}
