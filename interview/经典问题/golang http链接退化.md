在 HTTP 1.1 协议中，客户端和服务器端可以通过 Connection 首部字段来指定使用的连接方式，包括 Keep-Alive 和 Close 两种。Keep-Alive 意味着使用长连接，而 Close 意味着使用短连接。如果服务器端在处理请求时遇到性能问题，就可能会主动关闭连接，从而导致客户端降级为 HTTP 1.0 协议的短连接方式。

以下是一个具体的例子，来说明如何将 HTTP 1.1 降级为 HTTP 1.0：

假设我们有一个使用 HTTP 1.1 协议的网站，客户端通过浏览器访问该网站。浏览器首先会建立一条 HTTP 1.1 的连接到服务器端，然后发送多个请求。由于 HTTP 1.1 支持长连接，这些请求可以复用一个连接，从而提高性能和并发量。

但是，如果服务器端出现性能问题，或者在处理请求时遇到其他故障，就可能主动关闭连接，从而导致客户端降级为 HTTP 1.0 协议的短连接方式。在这种情况下，服务器会在响应报文中添加 Connection: Close 首部字段，告知客户端关闭连接，不再使用长连接。此时客户端将无法复用连接，每个请求都需要重新建立一条新的连接，这会导致性能下降，但是能够保证请求的可靠性。

总的来说，HTTP 1.1 协议的降级和升级都是比较容易发生的，这需要我们在实际开发中根据性能、安全、稳定性等因素来选择合适的协议和策略，以优化应用程序的性能和用户体验。

以下是使用调优后的 HTTP 客户端对 GitHub 发起请求的示例代码，其中使用优化后的连接池参数,
```go
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)

func main() {
    // 设置连接池参数
    http.DefaultTransport.(*http.Transport).MaxIdleConns = 200
    http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 100
    http.DefaultTransport.(*http.Transport).IdleConnTimeout = 120 * time.Second
    http.DefaultTransport.(*http.Transport).TLSHandshakeTimeout = 15 * time.Second

    // 构造请求
    req, err := http.NewRequest("GET", "https://api.github.com", nil)
    if err != nil {
        fmt.Println("Error creating request:", err)
        return
    }

    // 添加请求头
    req.Header.Set("User-Agent", "My Go HTTP Client")

    // 发起请求
    client := http.DefaultClient
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return
    }

    // 处理响应
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response:", err)
        return
    }
    fmt.Println(string(body))
}
```

源码分析：
https://cloud.tencent.com/developer/article/1900690
https://zhuanlan.zhihu.com/p/451642373
https://studygolang.com/articles/12720
https://vearne.cc/archives/39913
https://www.cnblogs.com/JulianHuang/p/15950624.html