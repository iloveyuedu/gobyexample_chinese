// [regular expressions](http://en.wikipedia.org/wiki/Regular_expression).
// go为正则表达式提供了内建支持
// 下面是在go里使用正则表达式执行一些常见的任务

package main

import "bytes"
import "fmt"
import "regexp"

func main() {

    // 测试一个正则是否匹配一个字符串
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)

    // 上面我们直接使用了一个正则表达式字符串，但是一般我们
    // 都需要先编译正则表达式来优化正则表达式的结构
    r, _ := regexp.Compile("p([a-z]+)ch")

    // 正则表达式有许多方法可以使用
    // 下面就是上面匹配字符串更好的写法
    fmt.Println(r.MatchString("peach"))

    // 用正则去查找字符串
    fmt.Println(r.FindString("peach punch"))

    // 用正则去查找字符串，返回匹配的起始和结束索引
    fmt.Println(r.FindStringIndex("peach punch"))

    // `子匹配`比全匹配提供了更多的信息，比如下面这个子匹配
    // 将同时返回`p([a-z]+)ch` 和 `([a-z]+)` 的匹配信息
    fmt.Println(r.FindStringSubmatch("peach punch"))

    // 子匹配，并返回匹配到的字符串起始索引
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))

    // `All` 修饰的函数会从字符串里找到所有的匹配，而不仅仅像上面
    // 那样只匹配到一个就返回，比如就是查找所有的匹配
    fmt.Println(r.FindAllString("peach punch pinch", -1))

    // 其它的函数也可以使用`All` 来修饰
    fmt.Println(r.FindAllStringSubmatchIndex(
        "peach punch pinch", -1))

    // 提供第二个参数来指定匹配的次数
    fmt.Println(r.FindAllString("peach punch pinch", 2))

    // 上面的方法我们都使用字符串作为参数，其函数名都类似
    // `MatchString` 带有string字样，其实我们也可以使用
    // `[]byte`类型作为参数，去掉函数名里的string即可
    fmt.Println(r.Match([]byte("peach")))

    // 当我们创建一个正则表达式时使用`Compile`
    // 如果我们要确保正则表达式没问题就应该使用`MustCompile`
    // 如果正则表达式有问题将直接报错提示
    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println(r)

    // 在某些操作上 `regexp`包可以替代strings包
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

    // 有趣的是，你可以传递一个函数来修改匹配的东西:)
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}
