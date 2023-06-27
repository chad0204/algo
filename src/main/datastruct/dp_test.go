package datastruct

import (
	"fmt"
	"math"
	"testing"
)

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
		tmp = Min(tmp, v+1) // 少了c 多一枚硬币
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
			dp[i] = Min(dp[i-c]+1, dp[i])
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

/**

#状态
1 持有股票 0 未持有股票

当前天
0 <= i <= n-1

交易次数
1 <= j <= k

#选择
buy sell rest(无操作)


穷举
dp[i][j][0 or 1]
	for 0 <= i < n:
		for 1<= j <= k:
			for s in{0, 1}:
				dp[i][j][s] = max(buy, sell, rest)

dp[n-1][k][0]



#今天(i, 从0开始)没持有, 今天的交易笔数为j, dp[i][j][1]:
1. 昨天也没持有, 今天rest, dp[i-1][j][0]
2. 昨天持有, 今天卖了一笔, dp[i-1][j][1] + price[i]
dp[i][j][0] = max{dp[i-1][j][0],  dp[i-1][j][1] + price[i]}

#今天持有, 今天的交易笔数为j, dp[i][j][1]:
1. 昨天持有, 今天rest, dp[i-1][j][1]
2. 昨天没持有, 今天买了一笔, dp[i-1][j-1][0] - price[i]
dp[i][j][1] = max{dp[i-1][j][1],  dp[i-1][j-1][0] - price[i]}


base case:
i == -1, 第一天之前
dp[-1][x][0] = 0
dp[-1][x][1] = ? 不存在 负无穷 ???

j == 0, 一次交易没有
dp[x][0][1] = ? 不存在 负无穷
dp[x][0][0] = 0


题目的共同约束:
都是不能重复买或卖, 买和卖交替执行, 才算一笔完整交易。
1. ”你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票“
2. ”你在任何时候 最多只能持有一股股票“
3. ”你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票“

条件3不仅要求交替买卖, 还限制只能进行一次买卖。条件2说明可以无限次买卖。但是都要求交替买卖, 否则题目就不成立。




*/
func maxProfitK(k int, prices []int) int {
	n := len(prices)
	dpTable := make([][][]int, n)
	for i := range dpTable {
		dpTable[i] = make([][]int, k+1)
		for j := range dpTable[i] {
			dpTable[i][j] = make([]int, 2)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= k; j++ {
			if i == 0 {
				dpTable[i][j][0] = Max(0, math.MinInt32)
				dpTable[i][j][1] = Max(math.MinInt32, 0-prices[i])
				continue
			}
			if j == 0 {
				dpTable[i][j][0] = 0
				dpTable[i][j][1] = math.MinInt32
				continue
			}
			dpTable[i][j][0] = Max(dpTable[i-1][j][0], dpTable[i-1][j][1]+prices[i])
			dpTable[i][j][1] = Max(dpTable[i-1][j][1], dpTable[i-1][j-1][0]-prices[i])
		}
	}
	return dpTable[n-1][k][0]
}

// k =1
func maxProfit1(prices []int) int {
	/*
	   state func:
	   昨天没有, 今天无操作; 昨天有, 今天卖了。今天卖了不能买, 无妨。
	   dp[i][1][0] = max{dp[i-1][1][0], dp[i-1][1][1] + prices[i]}

	   昨天有, 今天无操作; 昨天没有, 今天买了。今天买了不能卖, 而且只能交易一次, 所以今天利润是-prices[i]
	   dp[i][1][1] = max{dp[i-1][1][1], dp[i-1][0][0] - prices[i]}
	               = max{dp[i-1][1][1], -prices[i]}

	   base case:
	   dp[-1][x][0] = 0
	   dp[-1][x][1] = math.MinInt32

	   dp[i][0][0] = 0
	   dp[j][0][1] = math.MinInt32

	*/
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[0][0] = 0
			dp[0][1] = -prices[i]
			continue
		}
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][1], -prices[i])
	}

	return dp[n-1][0]
}

// k = +infinity
func maxProfitInfinity(prices []int) int {
	/*
	   state func:
	   dp[i][0] = max{dp[i-1][0], dp[i-1][1] + prices[i]}
	   dp[i][1] = max{dp[i-1][1], dp[i-1][0] - prices[i]}

	   base case:
	   dp[-1][0] = 0
	   dp[-1][1] = math.MinInt32
	*/
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[0][0] = 0
			dp[0][1] = -prices[i]
			continue
		}
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}
