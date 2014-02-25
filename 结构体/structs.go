// go的结构体是多个字段的集合，在为数据分组时显得非常
// 有用

package main

import "fmt"

// `person`结构体有 `name` 和 `age` 两个字段
type person struct {
    name string
    age  int
}

func main() {

    // 下面这种语法创建一个新的struct
    fmt.Println(person{"Bob", 20})

    // 在初始化struct可以指定字段名字初始化
    fmt.Println(person{name: "Alice", age: 30})

    // 没有初始化的字段将得到那个字段的 "0值"
    fmt.Println(person{name: "Fred"})

    // `&`前缀表示这是一个指向struct的指针
    fmt.Println(&person{name: "Ann", age: 40})

    // 通过 . 来获取struct的字段值
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    // 你也可以在struct指针上使用.获取实际的值，
    // 指针会自动转为实际值
    sp := &s
    fmt.Println(sp.age)

    // Structs是可以被改变的
    sp.age = 51
    fmt.Println(sp.age)
}
