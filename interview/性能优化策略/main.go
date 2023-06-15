package main

import (
	"fmt"
	"math/big"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(4)

	// 这里是启动一个pprof的Http服务，监听/debug/pprof 路径下的外部指令
	go func() {
		defer wg.Done()
		fmt.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	// 以下为业务逻辑

	time.Sleep(time.Second * 5)

	go func() {
		defer wg.Done()
		var limitChan = make(chan int, 50)

		for true {
			limitChan <- 1
			go func() {
				time.Sleep(time.Millisecond * 100)
				<-limitChan
			}()
		}
	}()

	// 申请内存对象
	go func() {
		defer wg.Done()
		for true {
			testMap := make([]int64, 100000)
			fmt.Printf("容量=%v\n", len(testMap))
			time.Sleep(time.Millisecond * 100)
		}
	}()

	// Cpu密集型计算任务 计算pi
	go func() {
		defer wg.Done()
		for true {
			var digits = 1000000
			// 计算圆周率pai
			const blockSize = 100
			const radix = 10

			k := 0
			pi := new(big.Rat).SetInt64(0)

			for {
				if k > digits/blockSize {
					break
				}

				a := new(big.Int).Mul(big.NewInt(int64(blockSize)), big.NewInt(int64(2*k+1)))
				b := new(big.Int).Exp(big.NewInt(-1), big.NewInt(int64(k)), nil)
				c := new(big.Rat).SetFrac(b, a)

				pi.Add(pi, c)
				k++
			}

			pi.Mul(pi, big.NewRat(radix, 1))

			fmt.Println("pi:", pi)
			time.Sleep(time.Millisecond * 10) // Sleep 10毫秒
		}
	}()

	wg.Wait()
}

func calculatePi(digits int) *big.Rat {
	const blockSize = 100
	const radix = 10

	k := 0
	pi := new(big.Rat).SetInt64(0)

	for {
		if k > digits/blockSize {
			break
		}

		a := new(big.Int).Mul(big.NewInt(int64(blockSize)), big.NewInt(int64(2*k+1)))
		b := new(big.Int).Exp(big.NewInt(-1), big.NewInt(int64(k)), nil)
		c := new(big.Rat).SetFrac(b, a)

		pi.Add(pi, c)
		k++
	}

	pi.Mul(pi, big.NewRat(radix, 1))

	return pi
}
