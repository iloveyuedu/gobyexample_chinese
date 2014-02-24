// 在这个例子里，我们将用协程和通道实现工作池

package main

import "fmt"
import "time"

// 这是工作函数，我们将用它作为几个并发的实例
// 这个并发的实例将会通过jobs通道发送响应结果，每个
// 工作执行之间我们都会暂停几秒来模拟执行起来“昂贵”的任务
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "processing job", j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func main() {

    // 为了使用使用我们的工作池，我们需要给工作函数发送工作内容
    // 并收集它们的执行结果，因此我们申明了两个通道
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // 下面启动了三个工作函数，初始是阻塞的，因为
    // 它们还没有工作,jobs通道内没内容:)
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // 下面我们发送了九个工作内容到jobs通道内，然后我们
    // 关闭了jobs通道来表明我们这就是我们所有要执行的工作
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    // 最后我们收集所有工作函数的执行结果
    for a := 1; a <= 9; a++ {
        <-results
    }
}
