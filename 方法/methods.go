// Go 支持在struct类型上定义方法.

package main

import "fmt"

type rect struct {
    width, height int
}

// `area`方法绑定到`*rect`类型上
func (r *rect) area() int {
    return r.width * r.height
}

// 方法可以绑定到结构体指针和结构体值上，下面我们看看
// 如何将方法绑定到结构体值上
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{width: 10, height: 5}

    // 下面我们调用了刚定义好的两个方法
    fmt.Println("area: ", r.area())
    fmt.Println("perim:", r.perim())

    // go会为方法调用自动的处理值和指针,你可能想要使用指针
    // 来接收参数避免发生结构体的复制，可以像下面直接用指针进行操作
    rp := &r
    fmt.Println("area: ", rp.area())
    fmt.Println("perim:", rp.perim())
}
