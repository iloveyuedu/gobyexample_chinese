// 在程序里会经常获取unix计时到现在的 秒，微秒，纳秒的操作
// [Unix epoch](http://en.wikipedia.org/wiki/Unix_time).
// 下面是在go里面进行操作

package main

import "fmt"
import "time"

func main() {

    // 使用`time.Now`的结果获取从unix时间到现在逝去的秒和纳秒(`Unix` or `UnixNano`)
    now := time.Now()
    secs := now.Unix()
    nanos := now.UnixNano()
    fmt.Println(now)

    // 注意这里是不能直接获取微秒的，所以你需要用纳秒除以
	// 微秒的换算单位来获取微秒
    millis := nanos / 1000000
    fmt.Println(secs)
    fmt.Println(millis)
    fmt.Println(nanos)

    // 你可以将整形的秒或纳秒转成go的时间对象
    fmt.Println(time.Unix(secs, 0))
    fmt.Println(time.Unix(0, nanos))
}
