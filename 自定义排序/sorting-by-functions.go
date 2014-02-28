// 有时候我们想自定义去排序集合变量，比如我们想按
// 字符串的长度进行排序而不是按照字母顺序去排，下面
// 使用go进行自定义排序

package main

import "sort"
import "fmt"

// 我们需要一个自定义类型来响应go的排序函数，来实现自定义排序
// 下面我们创建了一个`ByLength`方法，它仅仅是内建`[]string`类型的
// 别名
type ByLength []string

// 我们需要实现 `sort.Interface`接口的 `Len`, `Less`,
// `Swap`方法，这样我们才能使用sort包的`Sort`函数进行自定义排序
// `Len` 和 `Swap`在实际的排序算法里会用到，`Less`用于比较两个元素
// 我们的例子里我们想要以字符串长度递增来排序字符串切片，所以我们下面使用
// 了`len(s[i])` 和 `len(s[j])`
func (s ByLength) Len() int {
    return len(s)
}
func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

// 实现的所有代码就下面这些[简单吧?],我们需要把原始的`fruits`
// 字符串切片转成`ByLength`类型，这个实现了排序包接口的类型
// 然后我们就可以使用`sort.Sort`进行自定义排序了
func main() {
    fruits := []string{"peach", "banana", "kiwi"}
    sort.Sort(ByLength(fruits))
    fmt.Println(fruits)
}
