// 经常用来为二进制或问题计算短标识符,比如[git 版本控制系统](http://git-scm.com/)
// 使用SHA1s来表示各个版本文件和目录，下面是在go里计算SHA1哈希

package main

// `crypto/*`包里实现了多个哈希函数
import "crypto/sha1"
import "fmt"

func main() {
    s := "sha1 this string"

    // `sha1.New()`,`sha1.Write(bytes)`, `sha1.Sum([]byte{})`是生成sha1哈希的模式
    // 下面我们就来生成一个哈希
    h := sha1.New()

    // `Write`期望一个bytes类型，所以我们需要将字符串强制转成bytes类型
    h.Write([]byte(s))

    // 下面获取最终的哈希值，是以byte切片类型的结果，`Sum`的参数将会
    // 被额外的加入到已存在的byte切片内，这通常是不需要的
    bs := h.Sum(nil)

    // SHA1值通常用16进制打印，比如在git提交时，使用`%x`将
    // 结果转为16进制字符串
    fmt.Println(s)
    fmt.Printf("%x\n", bs)
}
