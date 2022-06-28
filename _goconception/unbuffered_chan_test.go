package _goconception

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

var (
	wgp sync.WaitGroup
)

// 参考Go语言实战
func PlayCourt() {
	wgp.Add(2)
	court := make(chan int)
	rand.Seed(time.Now().UnixNano())
	go Player("Nadal", court)
	go Player("Djokovic", court)
	court <- 1
	wgp.Wait()
}

func Player(player string, court chan int) {
	defer wgp.Done()
	for {
		ball, ok := <-court
		if !ok {
			fmt.Printf("%s win\n", player)
			return
		}

		// 随机判断是否可以接住
		num := rand.Intn(100)
		if num%7 == 0 {
			fmt.Printf("%s defeat\n", player)
			close(court)
			return
		}

		fmt.Printf("%s hit ball %d\n", player, ball)
		ball++
		court <- ball
	}
}

func OverLoadChan() {
	var m = make(chan int, 10)
	for i := 0; i < 100; i++ {
		m <- i
		// 当写入的数据超过通道的长度，程序会被阻塞住
	}
}

func TestChan(t *testing.T) {
	PlayCourt()
	//OverLoadChan()
}
