package main

import "github.com/fzzbt/radix/redis"
import "time"
import "fmt"

func main() {
    // 初始化
    conf := redis.DefaultConfig()
    client := redis.NewClient(conf)
    fmt.Println(conf)
    fmt.Println(client)

    // set & get操作
    setRep := client.Set("foo", "bar")
    if setRep.Err != nil {
        panic(setRep.Err)
    }
    fmt.Println(setRep)
    getRep := client.Get("foo")
    if getRep.Err != nil {
        panic(getRep.Err)
    }
    getStr, _ := getRep.Str()
    fmt.Println(getRep)
    fmt.Println(getStr)

    // 同时获取不定个数的键Mget,使用了可变参数
    client.Set("foo1", "bar1")
    client.Set("foo2", "bar2")
    client.Set("foo3", "bar3")
    mgetRep := client.Mget("foo1", "foo2", "foo3")
    fmt.Println(mgetRep)

    // 调用多次
    mcRep := client.MultiCall(func(mc *redis.MultiCall) {
        mc.Set("k1", "v1")
        mc.Get("k1")
    })
    if mcRep.Err != nil {
        panic(mcRep.Err)
    }
    mcVal, _ := mcRep.Elems[1].Str()
    fmt.Println(mcVal)

    // 事务操作
    tRep := client.Transaction(func(mc *redis.MultiCall) {
        mc.Set("k2", "v2")
        mc.Get("k2")
    })
    if tRep.Err != nil {
        panic(tRep.Err)
    }
    tStr, _ := tRep.Elems[1].Str()
    fmt.Println(tStr)

    // 订阅/发布
    msgHdlr := func(msg *redis.Message) {
        fmt.Println(msg)
    }
    sub, subErr := client.Subscription(msgHdlr)
    if subErr != nil {
        panic(subErr)
    }
    defer sub.Close()
    sub.Subscribe("chan1", "chan2")
    sub.Psubscribe("chan*")
    client.Publish("chan1", "foo")
    sub.Unsubscribe("chan1")
    client.Publish("chan2", "bar")
    time.Sleep(time.Second)
}

// todo: 连接池 & 并发?
// todo: 重新连接?
// todo: 错误?
// todo: redis_url
