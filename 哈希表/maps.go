// 哈希表是go的内建类型[associative data type](http://en.wikipedia.org/wiki/Associative_array)
// 在别的语言里可能叫做字典(python),hash表(java)

package main

import "fmt"

func main() {

	// 使用内建的 `make`方法创建一张哈希表
    // `make(map[key-type]val-type)`.
    m := make(map[string]int)

	// 设置值使用熟悉的 `name[key] = val`语法
    m["k1"] = 7
    m["k2"] = 13

	// 通过 `Println` 打印将输出该哈希表所有的键值对
    fmt.Println("map:", m)

	// 通过 `name[key]`获取一个值
    v1 := m["k1"]
    fmt.Println("v1: ", v1)

	// 内建的 `len` 函数获取哈希表里键值对的个数
    fmt.Println("len:", len(m))

	// 内建的 `delete` 函数用来删除指定的键值对
    delete(m, "k2")
    fmt.Println("map:", m)

    // 可选接收的第二个返回值可用来判断当前的键是否存在于
	// 该张哈希表中, 这样可以区别返回值，因为实际情况下一个
	// 键可能对应的是0值(布尔为false)
    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    // 你可以用下面的方式来同时申明和初始化一张哈希表
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}
