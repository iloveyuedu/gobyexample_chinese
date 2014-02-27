// go里使用明确的方式来处理错误，将错误作为返回值返回(go里可以有多个返回值)
// 在其它语言里比如java和ruby会将错误作为函数结果返回(只能返回一个值)
// 有时候也可以在c语言里看到,go会使处理错误更简单些，没有出错也更明白些

package main

import "errors"
import "fmt"

// 惯例上，错误是作为函数的最后一个返回值，并且类型为 `error`(一个内建的接口)
func f1(arg int) (int, error) {
    if arg == 42 {

        // `errors.New` 通过错误信息构建了一个基本的 `error` 类型值
        return -1, errors.New("can't work with 42")

    }

    // 错误为nil表示没有出现错误
    return arg + 3, nil
}

// 尽可能自己去实现 `Error()` 接口的方法，下面的argError
// 就是一个自定义实现,代表参数错误
type argError struct {
    arg  int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 42 {

        // 这里我们使用 `&argError` 建立一个新的结构体，并
        // 给两个字段都赋了值
        return -1, &argError{arg, "can't work with it"}
    }
    return arg + 3, nil
}

func main() {

    // 下面两个循环用来测试出现错误的返回值，注意
    // 这里在if里内联进行了错误处理,这在go里是常见的
    // 处理方式
    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed:", e)
        } else {
            fmt.Println("f1 worked:", r)
        }
    }
    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 failed:", e)
        } else {
            fmt.Println("f2 worked:", r)
        }
    }

    // 如果你想要在程序里使用错误数据，你需要通过类型断言
    // 获取自定义错误的实例
    _, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}
