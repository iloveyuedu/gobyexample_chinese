// 在go里，数组是一个固定长度的有序元素的序列

package main

import "fmt"

func main() {

    // 我们创建了一个名为 a 的数组，它的长度为5的，类型为
    // int的数组,元素的类型和数组的长度都是这个数组的一部
    // 分，新创建的数组将默认已经"0值初始化",对于int类型的
    // 数组意味着这个数组的所有元素初始值为0
    var a [5]int
    fmt.Println("emp:", a)

    // 我们可以通过索引去 设置/获取 元素的值
    // `array[index] = value` 设置指定索引的值
    // `array[index]` 获取指定索引的值
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    // 通过 内建函数 `len` 获取数组的长度
    fmt.Println("len:", len(a))
    
    // 使用这种语法在一行里就可以申明一个数组并同时初始化
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    // 数组类型是一维的，但你可以通过组合类型
    // 来建立多维数组
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
