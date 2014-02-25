// go 的 `math/rand` 包提供了
// [伪随机数](http://en.wikipedia.org/wiki/Pseudorandom_number_generator)
// 生成器

package main

import "fmt"
import "math/rand"

func main() {

    // 例如下面的 `rand.Intn` 返回一个随机数整数n `0 <= n < 100`
    fmt.Print(rand.Intn(100), ",")
    fmt.Print(rand.Intn(100))
    fmt.Println()

    // `rand.Float64` 返回一个随机64位的浮点数`f`,
    // `0.0 <= f < 1.0`.
    fmt.Println(rand.Float64())

    // 生成一个指定范围内的随机浮点数，比如
    // `5.0 <= f' < 10.0`.
    fmt.Print((rand.Float64()*5)+5, ",")
    fmt.Print((rand.Float64() * 5) + 5)
    fmt.Println()

    // 通过确定随机数种子来控制随机数行为
    s1 := rand.NewSource(42)
    r1 := rand.New(s1)

    // 调用`rand.Source`和 `rand` 包函数调用一样
    fmt.Print(r1.Intn(100), ",")
    fmt.Print(r1.Intn(100))
    fmt.Println()

    // 如果种子和上面一样，你将得到同样的随机结果序列
    s2 := rand.NewSource(42)
    r2 := rand.New(s2)
    fmt.Print(r2.Intn(100), ",")
    fmt.Print(r2.Intn(100))
    fmt.Println()
}
