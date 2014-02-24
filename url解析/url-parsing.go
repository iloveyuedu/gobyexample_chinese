// URLs 提供了一个 [正式的方式定位资源](http://adam.heroku.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/).
// 下面是如何使用go解析url
package main

import "fmt"
import "net/url"
import "strings"

func main() {

    // 我们解析的这个url，包含了 协议，授权信息，域名，端口，路径
    // 和查询字符串
    s := "postgres://user:pass@host.com:5432/path?k=v#f"

    // 解析这个URL并确定没有发生错误
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    // 获取URL使用的协议
    fmt.Println(u.Scheme)

    // `User` 字段包含了所有的授权信息，`Username` 和 `Password`
    // 提供了额外的信息
    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)

    // `Host` 包含了域名和端口[存在的话],`Split`方法作用`Host`
    // 可以分离出域名和端口
    fmt.Println(u.Host)
    h := strings.Split(u.Host, ":")
    fmt.Println(h[0])
    fmt.Println(h[1])

    // 下面我们获取`path`和`#`符之后的内容
    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

    // 将查询字符串获取成 `k=v` 格式使用 `RawQuery`,也可以将
    // 查询字符串解析到哈希表内，哈希表内的查询值保存在一个
    // 字符串切片内，所以索引`[0]`将获取到第一个值
    fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])
}
