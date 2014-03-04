// 在之前的例子里我们了解到了如何使用atomic管理计数器的状态
// 更复杂的状态我们可以使用_[锁](http://en.wikipedia.org/wiki/Mutual_exclusion)_
// 来安全的获取多个协程共同操作的数据

package main

import (
    "fmt"
    "math/rand"
    "runtime"
    "sync"
    "sync/atomic"
    "time"
)

func main() {

    // `state`定义为一张表
    var state = make(map[int]int)

    // `mutex`将会排队去操作`state`
    var mutex = &sync.Mutex{}

    // 和我们之前看过的锁机制进行比较，`ops`会统计
    // 我们对state进行了多少次操作
    var ops int64 = 0

    // 下面我们启动了一百个协程用来执行对state的读操作
    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {

                // 每次我们都随机出一个键来从state表中取数据
                // 同时我们使用`mutex`里的`Lock()`来确保在对
                // state操作时，state被当前协程独占的
                // 获取完数据后执行`Unlock()`来释放锁
                // 同时对`ops`执行加一操作
                key := rand.Intn(5)
                mutex.Lock()
                total += state[key]
                mutex.Unlock()
                atomic.AddInt64(&ops, 1)

                // In order to ensure that this goroutine
                // doesn't starve the scheduler, we explicitly
                // yield after each operation with
                // `runtime.Gosched()`. This yielding is
                // handled automatically with e.g. every
                // channel operation and for blocking
                // calls like `time.Sleep`, but in this
                // case we need to do it manually.
                // 为了确保当前这个协程也能被调度，这里我们使用`runtime.Gosched()`
                // 来让出时间片
                runtime.Gosched()
            }
        }()
    }

    // 我们使用和读操作相同的模式，开10个协程来模拟写操作
    for w := 0; w < 10; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddInt64(&ops, 1)
                runtime.Gosched()
            }
        }()
    }

    // 让上面的10个协程先跑一会儿
    time.Sleep(time.Second)

    // 获取并输出最后的计数结果
    opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:", opsFinal)

    // With a final lock of `state`, show how it ended up.
    // 最后锁一下`state`，看看最终是什么情况
    mutex.Lock()
    fmt.Println("state:", state)
    mutex.Unlock()
}
