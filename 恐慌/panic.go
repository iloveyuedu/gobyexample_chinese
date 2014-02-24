// 一个 `panic` 通常意味着某样东西出现了错误，大多数情况
// 我们应该使用它来找到错误，而不应该出现在普通操作中，或者
// 我们还没有准备好如何优雅的去处理它

package main

import "os"

func main() {

    // 我们通常会抛出panic来表明出现了非预期的错误，这通常
    // 是唯一使用panic的地方
    panic("a problem")

    // 通常一个函数返回值出现了错误,而我们又不想处理它，
    // 那么我们将使用panic来取消后面的操作,如下面这个例子
    // 当我们创建一个文件出现错误时使用panic抛出错误
    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}
