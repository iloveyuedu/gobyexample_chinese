// 之前的关于range的例子中我们知道如何结合使用`for` 和 `range`
// 来遍历基础集合类型数据结构,我们也可以使用这种语法去遍历一个
// 等待输出数据的通道

package main

import "fmt"

func main() {

    // We'll iterate over 2 values in the `queue` channel.
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    // `range`会遍历每一个 `queue` 通道收到的元素，因为
    // 我们在上面已经关闭了 `queue` 通道，下面的遍历会在
    // 获取到两个元素后结束，如果我们上面没有关闭 `queue` 通道
    // 那么下面这个循环将会阻塞，以等待接收第三个元素
    for elem := range queue {
        fmt.Println(elem)
    }
}
