package main

import (
	"fmt"
	"sync"
)

// main 编译会出现并发写的错误
func main() {
	var wg sync.WaitGroup
	m := make(map[int]int)
	// 优化
	var mu sync.RWMutex

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				mu.Lock() // 不加该锁 fatal error: concurrent map writes
				m[j] = j
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	fmt.Println(m)
	fmt.Println(len(m))
}
