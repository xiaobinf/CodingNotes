在 Go 中，可以使用 time.NewTicker 方法创建一个定时器，以便在指定的时间间隔内执行某个操作。定时器会定期触发一个事件，并在每次触发事件时执行指定的操作。

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // 创建一个定时器，每隔 1 秒触发一次事件
    pingInterval := time.Second
    ticker := time.NewTicker(pingInterval)

    // 在定时器触发事件时执行操作
    go func() {
        for range ticker.C {
            fmt.Println("Ping")
        }
    }()

    // 等待一段时间后停止定时器
    time.Sleep(5 * time.Second)
    ticker.Stop()
    fmt.Println("Stopped")
}
```
在这段代码中，我们首先创建了一个定时器 ticker，并设置定时器的时间间隔为 1 秒。然后，我们使用 go 关键字创建一个新的 goroutine，在定时器触发事件时执行操作。在这个示例中，我们输出了一个字符串 "Ping"。

最后，我们使用 time.Sleep 方法等待一段时间后停止定时器，并输出一个字符串 "Stopped"。

需要注意的是，定时器会在程序退出前一直运行，直到被显式地停止。在使用定时器时，需要确保在适当的时候停止定时器，以避免资源泄漏。