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

    // The `Submatch` variants include information about
    // both the whole-pattern matches and the submatches
    // within those matches. For example this will return
    // information for both `p([a-z]+)ch` and `([a-z]+)`.
    fmt.Println(r.FindStringSubmatch("peach punch"))

    // Similarly this will return information about the
    // indexes of matches and submatches.
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))

    // The `All` variants of these functions apply to all
    // matches in the input, not just the first. For
    // example to find all matches for a regexp.
    fmt.Println(r.FindAllString("peach punch pinch", -1))

    // These `All` variants are available for the other
    // functions we saw above as well.
    fmt.Println(r.FindAllStringSubmatchIndex(
        "peach punch pinch", -1))

    // Providing a non-negative integer as the second
    // argument to these functions will limit the number
    // of matches.
    fmt.Println(r.FindAllString("peach punch pinch", 2))

    // Our examples above had string arguments and used
    // names like `MatchString`. We can also provide
    // `[]byte` arguments and drop `String` from the
    // function name.
    fmt.Println(r.Match([]byte("peach")))

    // When creating constants with regular expressions
    // you can use the `MustCompile` variation of
    // `Compile`. A plain `Compile` won't work for
    // constants because it has 2 return values.
    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println(r)

    // The `regexp` package can also be used to replace
    // subsets of strings with other values.
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

    // The `Func` variant allows you to transform matched
    // text with a given function.
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}
