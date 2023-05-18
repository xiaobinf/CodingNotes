```go
package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	m     map[string]string
	mutex sync.Mutex
}

func NewCache() *Cache {
	return &Cache{
		m: make(map[string]string),
	}
}

func (c *Cache) Set(key, value string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.m[key] = value
}

func (c *Cache) Get(key string) (string, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	value, ok := c.m[key]
	return value, ok
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.m, key)
}

func (c *Cache) Clean(interval time.Duration) {
	for {
		time.Sleep(interval)
		c.mutex.Lock()
		for key, value := range c.m {
			// 判断是否过期，如果过期则删除
			if isExpired(value) {
				delete(c.m, key)
			}
		}
		c.mutex.Unlock()
	}
}

func isExpired(value string) bool {
	// 判断是否过期，这里假设过期时间为 10 秒
	return time.Now().Unix()-parseTime(value) > 10
}

func parseTime(value string) int64 {
	// 解析时间戳
	t, _ := time.Parse(time.RFC3339, value)
	return t.Unix()
}

func main() {
	cache := NewCache()
	go cache.Clean(1 * time.Second)

	cache.Set("key1", "value1")
	cache.Set("key2", "value2")
	cache.Set("key3", "value3")

	time.Sleep(5 * time.Second)

	value, ok := cache.Get("key1")
	if ok {
		fmt.Println("key1:", value)
	} else {
		fmt.Println("key1 not found")
	}

	value, ok = cache.Get("key2")
	if ok {
		fmt.Println("key2:", value)
	} else {
		fmt.Println("key2 not found")
	}

	value, ok = cache.Get("key3")
	if ok {
		fmt.Println("key3:", value)
	} else {
		fmt.Println("key3 not found")
	}
}
```

在上面的代码中，我们定义了一个 `Cache` 结构体，其中包含一个 `map` 和一个互斥锁。`Set` 方法用于向 `map` 中添加键值对，`Get` 方法用于获取指定键的值，`Delete` 方法用于删除指定键的值。`Clean` 方法用于定期清理 `map` 中的无用数据，它会每隔一段时间遍历 `map` 中的所有键值对，判断是否过期，如果过期则删除。

在 `main` 函数中，我们创建了一个 `Cache` 实例，并调用 `Set` 方法向 `map` 中添加了三个键值对。然后，我们让程序休眠 5 秒钟，等待 `Clean` 方法清理无用数据。最后，我们调用 `Get` 方法获取三个键的值，判断它们是否存在。

需要注意的是，上面的代码只是一个简单的示例，实际应用中需要根据具体情况进行调整。例如，需要考虑并发访问的问题，以及如何优化清理算法等。