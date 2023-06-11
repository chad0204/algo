package datastruct

import (
	"fmt"
	"math"
	"testing"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestFib(t *testing.T) {
	fmt.Println(fib(1000))
}

// 斐波那契数列 dp(n) = dp(n-1) + dp(n-2)
func fib(n int) int {
	m := make(map[int]int)
	return f(n, m)
}

// 递归 自顶向下
func f(n int, m map[int]int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if m[n] != 0 {
		return m[n]
	}
	m[n] = f(n-1, m) + f(n-2, m)
	return m[n]
}

// 迭代 自底向上
func fibV2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	table := make([]int, n+1)
	table[0] = 0
	table[1] = 1

	for i := 2; i < n+1; i++ {
		table[i] = table[i-1] + table[i-2]
	}
	return table[n]
}

func TestCoinChange(t *testing.T) {
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
	fmt.Println(coinChangeIterator([]int{2, 5, 10, 1}, 27))
}

// 322. 零钱兑换 dp[amount] = min{dp[amount - coin] + 1}
func coinChange(coins []int, amount int) int {
	m := make([]int, amount+1)
	for i := range m {
		m[i] = -999
	}
	return dp(coins, amount, m)
}

func dp(coins []int, amount int, m []int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	if m[amount] != -999 { // 数组初始都设置成-999
		return m[amount]
	}

	tmp := math.MaxInt32
	for _, c := range coins {
		v := dp(coins, amount-c, m) //子问题
		if v == -1 {                // 子问题无解 父问题也无解
			continue
		}
		tmp = min(tmp, v+1) // 少了c 多一枚硬币
	}
	m[amount] = tmp
	if tmp == math.MaxInt32 {
		return -1
	}
	return tmp
}

func coinChangeIterator(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32 //dp[i-c]+1其他语言会溢出, 其实可以改成amount+1
	}
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		for _, c := range coins {
			if i < c {
				continue
			}
			dp[i] = min(dp[i-c]+1, dp[i])
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
