// [命令行标记](http://en.wikipedia.org/wiki/Command-line_interface#Command-line_option)
// 通常用来为命令行程序指定选项
// 比如在 `wc -l`命令中 这个 `-l` 就是一个命令行标记
package main

// go提供了一个 `flag` 包来支持基本的命令行标记解析，我们使用
// 这个包来实现我们的命令行程序
import "flag"
import "fmt"

func main() {

    // 可以定义标记为字符串，整形和布尔型，这里我们定义了一个
    // 字符串标记 `word` 且其默认值为`"foo"`并且有一个短的描述
    // `flag.String`函数返回一个字符串指针(不是一个字符串值)
    // 下面我们将看到如何使用这个字符串指针
    wordPtr := flag.String("word", "foo", "a string")

    // 定义名为 `numb` 和 `fork` 的标记，使用和上面相似的
    // 方法来获取其值
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    // 也可以使用一个存在的变量来作为标记定义的一部分
    // 更方便的获取命令行标记的值，注意我们需要传递指针到
    // StringVar函数里
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    // 当所有的标记都定义完成之后，通过调用`flag.Parse()`
    // 来完成命令行标记的解析
    flag.Parse()

    // 下面我们打印出所有解析之后的结果和命令行参数
    // 注意我们需要对指针进行解引用`*wordPtr`来获取其指向的值
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}
