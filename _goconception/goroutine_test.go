package _goconception

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TwoGoroutinePrint() {
	runtime.GOMAXPROCS(2)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for j := 0; j < 3; j++ {
			for i := 'a'; i < 'a'+26; i++ {
				fmt.Printf("%c\n", i)
			}
		}
	}()
	go func() {
		defer wg.Done()
		for j := 0; j < 3; j++ {
			for i := 'A'; i < 'A'+26; i++ {
				fmt.Printf("%c\n", i)
			}
		}
	}()

	fmt.Println("wait goroutine finish")
	wg.Wait()
	fmt.Println("main finish")
}

// LongTask测试一个需要长时间运行的goroutine
func LongTask() {
	runtime.GOMAXPROCS(1)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 'a'; i < 'a'+26; i++ {
			fmt.Printf("%c\n", i)
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 'A'; i < 'A'+26; i++ {
			fmt.Printf("%c\n", i)
			time.Sleep(10000)
		}
	}()

	fmt.Println("wait goroutine finish")
	wg.Wait()
	fmt.Println("main finish")
}

var (
	count int64 = 0
	wg          = &sync.WaitGroup{}
	mutex       = &sync.Mutex{}
)

func goroutineResource() {
	wg.Add(2)
	//go IncrCountUnSafe(1)
	//go IncrCountUnSafe(2)
	//go IncrCountAtomic(1)
	//go IncrCountAtomic(2)
	go IncrCountMutex(1)
	go IncrCountMutex(2)
	wg.Wait()
	fmt.Println(count)
}

func IncrCountUnSafe(id int) {
	var value = count
	defer wg.Done()
	runtime.Gosched()
	value = value + 3
	count = value
}

func IncrCountAtomic(id int) {
	// 原子函数加锁访问资源
	atomic.AddInt64(&count, 3)
	runtime.Gosched()
	defer wg.Done()
}

func IncrCountMutex(id int) {
	// 设置临界区
	mutex.Lock()
	var value = count
	runtime.Gosched()
	value++
	count = value
	mutex.Unlock()
	defer wg.Done()
}

func TestTwoGoroutinePrint(t *testing.T) {
	//TwoGoroutinePrint()
	//LongTask()
	goroutineResource()
}
