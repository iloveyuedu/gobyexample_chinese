// go标准库的 `strings` 包提供了大量常用的字符串操作相关的
// 函数，下面的例子将给你初步的映像

package main

import s "strings"
import "fmt"

// 我们给 `fmt.Println` 起个短名，下面会大量使用此函数
var p = fmt.Println

func main() {

    // 下面是一些`strings`包函数的使用例子，注意这些函数都是string包的
    // 而不是在字符串对象本身,那就意味着我们需要将字符串作为参数
    // 传递到这些函数里
    p("Contains:  ", s.Contains("test", "es"))
    p("Count:     ", s.Count("test", "t"))
    p("HasPrefix: ", s.HasPrefix("test", "te"))
    p("HasSuffix: ", s.HasSuffix("test", "st"))
    p("Index:     ", s.Index("test", "e"))
    p("Join:      ", s.Join([]string{"a", "b"}, "-"))
    p("Repeat:    ", s.Repeat("a", 5))
    p("Replace:   ", s.Replace("foo", "o", "0", -1))
    p("Replace:   ", s.Replace("foo", "o", "0", 1))
    p("Split:     ", s.Split("a-b-c-d-e", "-"))
    p("ToLower:   ", s.ToLower("TEST"))
    p("ToUpper:   ", s.ToUpper("test"))
    p()

    // 你可以从[`strings`](http://golang.org/pkg/strings/)找到string包更多的函数

    // 下面不是 `strings` 包的内容，但值得说一说，go提供了方便的
    // 机制获取字符串的长度和通过索引得到单个字符
    p("Len: ", len("hello"))
    p("Char:", "hello"[1])
}
