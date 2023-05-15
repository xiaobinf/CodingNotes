package coinChange_1014

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func coinChange(coins []int, amount int) int {
	sort.Ints(coins)

	if amount == 0 {
		return 0
	}

	if coins[0] > amount {
		return -1
	}

	if coins[0] < amount && len(coins) == 1 && amount%coins[0] != 0 {
		return -1
	}

	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		var m = math.MaxInt32
		for j := 0; j < len(coins); j++ {
			// fmt.Println(amount, coins[j])
			if i >= coins[j] {
				m = min(dp[i-coins[j]]+1, m)
			}

		}
		dp[i] = m
	}

	fmt.Println(dp)
	if dp[amount] > 1<<32 {
		return -1
	} else {
		return dp[amount]
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func TestName(t *testing.T) {
	a := []int{5, 3, 5, 7}
	sort.Ints(a)
	//fmt.Println(a)
	//coinChange([]int{3, 2, 5}, 11)
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
	//fmt.Println(4 % 2)
	fmt.Println(1<<31-1 == math.MaxInt32)
	fmt.Println(math.MaxInt32 + 1)
}
