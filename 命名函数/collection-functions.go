// 我们经常需要做很多数据处理的操作，比如查询出所有的商品
// 放入到另一新的类型里，就需要一个自定义的函数

// 在有些语言里(java?)内置了一种叫泛型的数据结构和相关的一些算法，go不支持
// 泛型，通常go是通过一系列函数来将数据处理成用户需要的类型

// 下面用字符串切片类型进行举例,你可以直接拿来使用
// 提示，有些情况下将这些函数定义到类型上可能会更方便
// 而不是创建一系列叫做助手的函数

package main

import "strings"
import "fmt"

// 返回字符串 `t` 的索引,没找到返回-1
func Index(vs []string, t string) int {
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

// 如果字符串t在字符串切片vs里存在返回ture，不存在返回false
func Include(vs []string, t string) bool {
    return Index(vs, t) >= 0
}

// 如果字符串切片vs中有元素满足断言函数`f`返回true，否者返回false
func Any(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}

// 如果字符串切片vs中所有元素都满足断言函数`f`返回true，否者返回false
func All(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if !f(v) {
            return false
        }
    }
    return true
}

// 返回一个满足断言函数 `f` 的所有子元素的新切片
func Filter(vs []string, f func(string) bool) []string {
    vsf := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

// 返回一个新的切片，该切片的元素是vs切片里经过函数f处理过的元素
func Map(vs []string, f func(string) string) []string {
    vsm := make([]string, len(vs))
    for i, v := range vs {
        vsm[i] = f(v)
    }
    return vsm
}

func main() {

    // 这里我们试试我们刚定义的几个集合函数
    var strs = []string{"peach", "apple", "pear", "plum"}

    fmt.Println(Index(strs, "pear"))

    fmt.Println(Include(strs, "grape"))

    fmt.Println(Any(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(All(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(Filter(strs, func(v string) bool {
        return strings.Contains(v, "e")
    }))
    
    // 上面的例子使用了匿名函数，你也可以使用正确类型的命名函数
    fmt.Println(Map(strs, strings.ToUpper))

}
