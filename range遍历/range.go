// 很多情况我们都会去遍历一个数据集合，下面让我们看看如何
// 使用`range`去遍历一个我们已知的数据结构

package main

import "fmt"

func main() {

    // 下面我们使用  `range` 遍历一个整形数组切片
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    // `range` 一个数组或数组切片，会同时得到数组的索引
    // 和当前元素的值，上面的例子我们没有使用索引，我们使用
    // `_`标识符忽略掉了该值,但有时候我们确实也需要使用该值
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    // `range`一个哈希表，并同时获得键和值
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    // `range` 一个字符串将会得到一个Unicode码点，第一个值
    // 是索引,第二个值是 `rune` 型的值
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
