// go 为json的编码/解码提供内建支持,同事支持
// 内建类型和自定义类型的编解码

package main

import "encoding/json"
import "fmt"
import "os"

// 我们将使用两种struct类型来展示如何编码
// 下面还有将json解码到自定义类型
type Response1 struct {
    Page   int
    Fruits []string
}
type Response2 struct {
    Page   int      `json:"page"`
    Fruits []string `json:"fruits"`
}

func main() {

    // 首先 我们看看如何将基础类型编码为json字符串，下面
    // 是编码原始值(值形式)
    bolB, _ := json.Marshal(true)
    fmt.Println(string(bolB))

    intB, _ := json.Marshal(1)
    fmt.Println(string(intB))

    fltB, _ := json.Marshal(2.34)
    fmt.Println(string(fltB))

    strB, _ := json.Marshal("gopher")
    fmt.Println(string(strB))

    // 下面是 切片和hash表 两种类型的编码，它会正确转成
    // 成json数组和对象两种形式
    slcD := []string{"apple", "peach", "pear"}
    slcB, _ := json.Marshal(slcD)
    fmt.Println(string(slcB))

    mapD := map[string]int{"apple": 5, "lettuce": 7}
    mapB, _ := json.Marshal(mapD)
    fmt.Println(string(mapB))

    // json 包会自动的编码自定义类型，它只会包含可输出的字段
    // 默认会使用字段名作为json的键
    res1D := &Response1{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res1B, _ := json.Marshal(res1D)
    fmt.Println(string(res1B))

    // 你也可以使用标记来自定义json键名
    // 检查下 `Response2` 类型的定义，该类型
    // 使用了自定义json键名
    res2D := &Response2{
        Page:   1,
        Fruits: []string{"apple", "peach", "pear"}}
    res2B, _ := json.Marshal(res2D)
    fmt.Println(string(res2B))

    // 现在，我们来看看如何将json数据解码到go变量里
    // 下面这个是一个普通的json字符串
    byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

    // 我们需要提供一个变量作为json能解码之后的容器,这个
    // `map[string]interface{}` 可以容纳任意类型的数据
    var dat map[string]interface{}

    // 下面是一个真实的解码实例，包含了相关错误的检查。
    if err := json.Unmarshal(byt, &dat); err != nil {
        panic(err)
    }
    fmt.Println(dat)

    // 为了使用解码之后的值，我们需要把它们转为正确的
    // 类型之后再使用,比如我们要将一个值转为我们期望的
    //  `float64` 类型并存储到num变量中。
    num := dat["num"].(float64)
    fmt.Println(num)

    // 获取嵌套的数据需要一点技巧
    strs := dat["strs"].([]interface{})
    str1 := strs[0].(string)
    fmt.Println(str1)

    // 我们也可以解码到自定义类型，这么做会让解码之后类型
    // 直接确定，不要要类型断言就可以直接使用
    str := `{"page": 1, "fruits": ["apple", "peach"]}`
    res := &Response2{}
    json.Unmarshal([]byte(str), &res)
    fmt.Println(res)
    fmt.Println(res.Fruits[0])

    // 在上面的例子中我们一直使用bytes和string类型作为中间体
    // 我们也可以直接编码到 `os.Writer`s 比如`os.Stdout`甚至
    // HTTP相应体中
    enc := json.NewEncoder(os.Stdout)
    d := map[string]int{"apple": 5, "lettuce": 7}
    enc.Encode(d)
}
