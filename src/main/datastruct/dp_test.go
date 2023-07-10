package datastruct

import (
	"fmt"
	"math"
	"testing"
)

/*

什么是动态规划:

最优子结构, 状态转移方程, base case, 状态, 选择


递归 + memory是本质, 迭代刷表是优化


思路：
先想递归
发现重复计算
通过记忆化等方法弄掉重复计算
最后看下能不能通过利用计算顺序来做到去掉递归用“刷表”方式直接顺序计算，能搞定最好搞不定拉倒

*/

func TestFib(t *testing.T) {
	fmt.Println(fib(1000))
}

// 斐波那契数列 dp(n) = dp(n-1) + dp(n-2)
func fib(n int) int {
	m := make(map[int]int)
	return fibV1(n, m)
}

// 递归 自顶向下
func fibV1(n int, m map[int]int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if m[n] != 0 {
		return m[n]
	}
	m[n] = fibV1(n-1, m) + fibV1(n-2, m)
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
	return dpCoin(coins, amount, m)
}

func dpCoin(coins []int, amount int, m []int) int {
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
		v := dpCoin(coins, amount-c, m) //子问题
		if v == -1 {                    // 子问题无解 父问题也无解
			continue
		}
		tmp = Min(tmp, v+1) // 少了c 多一枚硬币
	}
	m[amount] = tmp
	//这里也要注意
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

/*
股票买卖

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

题目的共同约束, 交替买卖:
都是不能重复买或卖, 买和卖交替执行, 才算一笔完整交易。
1. ”你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票“
2. ”你在任何时候 最多只能持有一股股票“
3. ”你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）“

条件1不仅要求交替买卖, 还限制只能进行一次买卖, 且必须是不同天。 121. 买卖股票的最佳时机
条件2没有说明次数, 可以无限次买卖, 但通过“只能持有一股“也要求交替买卖。 122. 买卖股票的最佳时机 II
条件3直接约束了不能交替买卖, 交易次数由题目给出, k次 or 2次。123. 买卖股票的最佳时机 III; 188. 买卖股票的最佳时机 IV

后面两种就是增加了不同的约束条件: 冷冻期, 手续费
*/
func maxProfitK(k int, prices []int) int {
	n := len(prices)
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= k; j++ {
			//i == 0 满足状态转移方程, 通过状态转移方程推倒出用base case表示。因为-1不能作为数组下标
			if i == 0 {
				dp[i][j][0] = Max(0, math.MinInt32)
				dp[i][j][1] = Max(math.MinInt32, 0-prices[i])
				continue
			}
			if j == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = math.MinInt32
				continue
			}
			dp[i][j][0] = Max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = Max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[n-1][k][0]
}

// k = 1
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

// 714. 买卖股票的最佳时机含手续费
func maxProfitWithFee(prices []int, fee int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			dp[0][0] = Max(0, math.MinInt32+prices[i])
			dp[0][1] = -prices[i] - fee
			continue
		}
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i]) // 也可以在卖的时候给手续费, 注意还是-fee
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
	}
	return dp[n-1][0]
}

// 309. 最佳买卖股票时机含冷冻期
func maxProfitWithFreeze(prices []int, fee int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	/*
	   sell之后是冷冻期。buy的前一天是冷冻期

	   dp[-1][0] = 0
	   dp[-1][1] = math.MinInt32
	   dp[-2][0] = 0
	   dp[-2][1] = math.MinInt32
	*/
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = Max(0, math.MinInt32)
			dp[i][1] = Max(math.MinInt32, -prices[i])
			continue
		}
		if i == 1 {
			dp[i][0] = Max(dp[0][0], dp[0][1]+prices[i])
			dp[i][1] = Max(dp[0][1], -prices[i])
			continue
		}
		//今天无: 昨天无; 昨天有今天卖了, 明天是冷冻期
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i])
		//今天有：昨天有; 昨天无,今天买了, 昨天是冷冻期, 状态转移是从前天而来。
		dp[i][1] = Max(dp[i-1][1], dp[i-2][0]-prices[i])
	}
	return dp[n-1][0]
}

/**

打家劫舍


https://mp.weixin.qq.com/s/z44hk0MW14_mAQd7988mfw

	不偷 0 偷 1

	这间不偷, 上一间偷了, 上一间没偷
	dp[i][0] = max(dp[i-1][1], dp[i-1][0])
	这间偷了, 上上间偷了, 上一间没偷
	dp[i][1] = max(dp[i-2][1] + nums[i], dp[i-1][0] + num[i])

	base case:
	dp[-1][0] = 0
	dp[-1][1] = math.MinInt32
	dp[-2][0] = 0
	dp[-2][1] = math.MinInt32


	可以使用一维数组和递归, 我的习惯是从0到n, 也可以从n到0


*/
// 198. 打家劫舍
func robV1(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = nums[i]
			continue
		}
		if i == 1 {
			dp[i][0] = Max(dp[0][1], dp[0][0])
			dp[i][1] = nums[i]
			continue
		}
		if i == 2 {
			dp[i][0] = Max(dp[1][1], dp[1][0])
			dp[i][1] = Max(dp[0][1]+nums[i], dp[1][0]+nums[i])
			continue
		}
		dp[i][0] = Max(dp[i-1][1], dp[i-1][0])
		dp[i][1] = Max(dp[i-2][1]+nums[i], dp[i-1][0]+nums[i])
	}
	return Max(dp[n-1][0], dp[n-1][1])
}

// 递归 自顶向下
func robV2(nums []int) int {
	mem := make([]int, len(nums))
	for i := range mem {
		mem[i] = -1
	}
	return dpRob(nums, len(nums)-1, mem)
}

func dpRob(nums []int, start int, mem []int) int {
	if start < 0 {
		return 0
	}
	if mem[start] != -1 {
		return mem[start]
	}
	res := Max(
		dpRob(nums, start-1, mem),             //start位置不抢,
		dpRob(nums, start-2, mem)+nums[start], //start位置抢, 只能去上上间
	)
	mem[start] = res
	return res
}

func TestRobV3(t *testing.T) {
	robV3([]int{2, 1})
}

// 一维数组 自低向上
func robV3(nums []int) int {

	n := len(nums)
	dp := make([]int, n)

	/*
		dp[-1] = 0
		dp[-2] = 0
	*/
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i] = nums[i]
			continue
		}
		if i == 1 {
			dp[i] = Max(dp[0], nums[i])
			continue
		}
		//i不抢, 和i-1一样; 今天抢了, i-2
		dp[i] = Max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

// 213. 打家劫舍 II
func robCycle(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return Max(robV3(nums[:len(nums)-1]), robV3(nums[1:]))
}

// 337. 打家劫舍 III
func robTree(root *TreeNode) int { // 还有思路是先层序后dp
	mem := make(map[*TreeNode]int)
	return robDp(root, mem)
}
func robDp(root *TreeNode, mem map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if _, ok := mem[root]; ok {
		return mem[root]
	}

	a := robDp(root.Left, mem) + robDp(root.Right, mem)

	b := root.Val // b表示抢, 所以加上val
	if root.Left != nil {
		b = b + robDp(root.Left.Left, mem) + robDp(root.Left.Right, mem)
	}
	if root.Right != nil {
		b = b + robDp(root.Right.Left, mem) + robDp(root.Right.Right, mem)
	}
	mem[root] = Max(a, b)
	return mem[root]
}
