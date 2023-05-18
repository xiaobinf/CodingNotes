package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Ban struct {
	visitIPs map[string]time.Time
	mu       sync.RWMutex
}

func NewBan() *Ban {
	return &Ban{visitIPs: make(map[string]time.Time)}
}
func (o *Ban) visit(ip string) bool {
	o.mu.Lock()
	defer o.mu.Unlock()
	_, ok := o.visitIPs[ip]
	if ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	return false
}
func main() {
	var success int64 = 0
	ban := NewBan()
	var wg sync.WaitGroup
	//var mu sync.RWMutex
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			wg.Add(1)
			go func(j int) {
				//fmt.Println(j)
				defer wg.Done()
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					atomic.AddInt64(&success, 1)
				}
			}(j)
		}
	}

	wg.Wait()
	fmt.Println("success:", success)
}
