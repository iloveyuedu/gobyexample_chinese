// go支持把固定的字符串解析成时间结构

package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    // 根据RFC3339格式化时间
    t := time.Now()
    p(t.Format("2006-01-02T15:04:05Z07:00"))

    // `Format` uses an example-based layout approach; it
    // takes a formatted version of the reference time
    // `Mon Jan 2 15:04:05 MST 2006` to determine the
    // general pattern with which to format the given
    // time. The example time must be exactly as shown: the
    // year 2006, 15 for the hour, Monday for the day of the
    // week, etc. Here are a few more examples of time
    // formatting.
    // 
    p(t.Format("3:04PM"))
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))

    // 对于存数字操作，你可以直接使用字符串格式化操作
    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

    // 输入的格式和输出的格式相同也是可以的
    withNanos := "2006-01-02T15:04:05.999999999-07:00"
    t1, e := time.Parse(
        withNanos,
        "2012-11-01T22:08:41.117442+00:00")
    p("t1",t1)
    kitchen := "3:04PM"
    t2, e := time.Parse(kitchen, "8:41PM")
    p(t2)

    // 如果输入字符串格式有误，你可以接收`Parse`函数的第二个返回
    // 值，来获取详细的出错信息，快速的定位问题
    ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    p(e)

    // 你也可以使用一些预定义的格式化格式,像下面这样
    p(t.Format(time.Kitchen))
}
