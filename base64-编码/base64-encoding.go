// go 为base64 编吗/解码 直接提供支持
// [base64encoding/decoding](http://en.wikipedia.org/wiki/Base64).

package main

// 这种语法是导入 `encoding/base64` 包，并重新起名为b64来代
// 替默认的 `base64`。这样有时候可能会方便一些
import b64 "encoding/base64"
import "fmt"

func main() {

	// 下面的 `字符串` 是我我们将要 编码/解码 的。
    data := "abc123!?$*&()'-=@~"

	// go 支持 标准和url兼容的base64编解码,
	// 下面是使用标准的编码器编码，编码器需要一个
	// `[]byte` 类型的变量，所以我们将把 `string` 转
	// 那种类型
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

	// 解码或许会返回一个错误，如果你不确定输入的格式是否有误，
	// 它会帮助你定位错误
    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

	// 使用 url兼容的 编解码器
    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}
