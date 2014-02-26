// 接口是方法集合的契约

package main

import "fmt"
import "math"

// 下面是为几何图形定义的通用的接口
type geometry interface {
    area() float64
    perim() float64
}

// 在我们的例子里，我们会在`square` 和 `circle`类型里实现上面那个接口
type square struct {
    width, height float64
}
type circle struct {
    radius float64
}

// 在go里实现一个接口，我们只需要实现接口定义的所有方法即可
// 下面为`square`s.类型实现`geometry`接口
func (s square) area() float64 {
    return s.width * s.height
}
func (s square) perim() float64 {
    return 2*s.width + 2*s.height
}

// `circle`s.类型实现`geometry`接口
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

// 如果一个变量是一个接口类型，那么这个变量可以接受所有实现
// 了那个接口的变量，而不管是什么类型，
// 下面这个函数就可以接收所有实现了`geometry`接口的变量
func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}

func main() {
    s := square{width: 3, height: 4}
    c := circle{radius: 5}

    // `circle` 和 `square`结构体都实现了`geometry`接口
    // 所以这两种类型的变量都可以作为measure函数的参数
    measure(s)
    measure(c)
}
