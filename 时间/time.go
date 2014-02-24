// go为时间和时间间隔操作提供了额外的支持
// 下面是几个实际的例子

package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    // 从获取当前时间开始
    now := time.Now()
    p(now)

    // You can build a `time` struct by providing the
    // year, month, day, etc. Times are always associated
    // with a `Location`, i.e. time zone.
    // 可以自定义一个时间结果来提供年月日天等,时间总是和地域
    // 关联的,比如timezone
    then := time.Date(
        2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
    p(then)

    // 从time变量中获取多个指定的时间信息
    p(then.Year())
    p(then.Month())
    p(then.Day())
    p(then.Hour())
    p(then.Minute())
    p(then.Second())
    p(then.Nanosecond())
    p(then.Location())

    // 星期几的信息也是可用的
    p(then.Weekday())

    // 下面这些方法比较两个时间，得到是否当前时间在前
    // 在后，同一时间
    p(then.Before(now))
    p(then.After(now))
    p(then.Equal(now))

    // `Sub` 方法返回两个时间间隔的 `Duration`表达式
    diff := now.Sub(then)
    p(diff)

    // 我们可以获取出两个时间间隔的长度
    p(diff.Hours())
    p(diff.Minutes())
    p(diff.Seconds())
    p(diff.Nanoseconds())

    // 你可以使用 `Add` 方法将时间向前推移, `-`将时间向后推移
    p(then.Add(diff))
    p(then.Add(-diff))
}
