package datastruct

import (
	"fmt"
	"math"
	"strings"
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
		m[i] = -1
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
	if m[amount] != -1 {
		return m[amount]
	}

	res := math.MaxInt32
	for _, c := range coins {
		v := dpCoin(coins, amount-c, m) //子问题
		if v == -1 {                    // 子问题无解 父问题也无解
			continue
		}
		res = Min(res, v+1) // 少了c 多一枚硬币
	}
	//res没有找到最小值，依然将math.MaxInt32设置到m[amount], 防止前面m[amount] == -1还要计算
	m[amount] = res
	//这里也要注意
	if res == math.MaxInt32 {
		return -1
	}
	return res
}

func coinChangeIterator(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		if i == 0 {
			dp[i] = 0
			continue
		}
		//初始化为math.MaxInt32, 可能dp[i-c]+1其他语言会溢出, 其实可以改成amount+1
		dp[i] = amount + 1
		for _, c := range coins {
			if i < c {
				continue
			}
			// 每次选一个硬币c, 那么就比i-c多一个硬币, 看下当前选哪个硬币最小
			dp[i] = Min(dp[i-c]+1, dp[i])
		}
	}
	if dp[amount] == amount+1 {
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
		dp[i] = make([][]int, k+1) // 因为从0开始, 比如两笔, 就有0, 1, 2
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
				dp[i][j][1] = math.MinInt32 //一笔没买就持有, 不可能啊。但是按照公式推应该是-prices[i]。都可以
				continue
			}
			dp[i][j][0] = Max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = Max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[n-1][k][0]
}

// k = 1 “今天买了不能卖” 和 k无限次的区别就是一天不能同时买卖
// 121. 买卖股票的最佳时机
func maxProfit1(prices []int) int {
	/*
	   state func:
	   昨天没有, 今天无操作; 昨天有, 今天卖了。今天卖了不能买, 无妨。
	   dp[i][j][0] = max{dp[i-1][j][0], dp[i-1][j][1] + prices[i]}

	   昨天有, 今天无操作; 昨天没有, 今天买了。今天买了不能卖, 而且只能交易一次, 所以今天利润是-prices[i]
	   dp[i][j][1] = max{dp[i-1][j][1], dp[i-1][j-1][0] - prices[i]}
	               = max{dp[i-1][1][1], -prices[i]}

	   去掉j

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
		dp[i][1] = Max(dp[i-1][1], -prices[i])
	}

	return dp[n-1][0]
}

// k = +infinity
// 122. 买卖股票的最佳时机 II
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
		dp[i][0] = Max(dp[i-1][1], dp[i-1][0]) // i-2
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

// 213. 打家劫舍 II 不要头的情况, 不要尾的情况
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

// 95. 不同的二叉搜索树 II
func numTreesDp(n int) int {
	mem := make([][]int, n+1)
	for i := range mem {
		mem[i] = make([]int, n+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	return numTreesDpHelper(1, n, mem)
}

func numTreesDpHelper(lo int, hi int, mem [][]int) int {
	if lo >= hi {
		return 1
	}
	if mem[lo][hi] != -1 {
		return mem[lo][hi]
	}
	res := 0
	for i := lo; i <= hi; i++ {
		l := numTreesDpHelper(lo, i-1, mem)
		r := numTreesDpHelper(i+1, hi, mem)
		res += l * r
	}
	mem[lo][hi] = res
	return res
}

// 300. 最长递增子序列
/*
思路：在i位置上lis等于任意比nums[i]小的元素的lis+1, 取最小值。dp[i] = min{dp[j]+1}, 0<=j<i,nums[j]<nums[i]
base case dp[0] = 1, 且任意位置的最小值都是1
*/
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		if i == 0 {
			dp[0] = 1
			continue
		}
		//遍历[0, i)得到dp[j], 找到比nums[i]大的nums[j]（说明dp[i]是dp[j]+1）。让dp[i]和这些dp[j]+1逐一比较
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
	}
	// 最后取最值
	res := 0
	for _, v := range dp {
		res = Max(v, res)
	}
	return res
}

// 931. 下降路径最小和. 思考： 是不是可以给长和宽都多一列(值为最大值, 这样就可以避免边界溢出)。试试自上而下的递归？
func minFallingPathSum(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i, _ := range dp {
		dp[i] = make([]int, len(matrix))
	}
	//行
	for i := 0; i < len(matrix); i++ {
		//列
		for j := 0; j < len(matrix); j++ {
			if i == 0 {
				dp[i][j] = matrix[i][j]
				continue
			}
			dp[i][j] = math.MaxInt32
			dp[i][j] = Min(dp[i-1][j]+matrix[i][j], dp[i][j])
			if j == 0 {
				dp[i][j] = Min(dp[i-1][j+1]+matrix[i][j], dp[i][j])
				continue
			}
			if j == len(matrix)-1 {
				dp[i][j] = Min(dp[i-1][j-1]+matrix[i][j], dp[i][j])
				continue
			}
			dp[i][j] = Min(dp[i-1][j+1]+matrix[i][j], dp[i][j])
			dp[i][j] = Min(dp[i-1][j-1]+matrix[i][j], dp[i][j])
		}
	}
	res := math.MaxInt32
	for i := 0; i < len(matrix); i++ {
		res = Min(dp[len(matrix)-1][i], res)
	}
	return res
}

// 思路 s[0: idx]在words中存在, 只需要判断s[idx:]是否存在即可, 逐步缩小范围, idx就是状态。减枝: 当words存在类似的单词, idx会被重复计算
func TestWord(t *testing.T) {

	wordBreak("leetcode", []string{"leet", "code"})
}

func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}
	//索引表示切割点, -1 未计算 0 已计算但凑不出 1 已计算能凑出
	memo := make([]int, len(s))
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}
	return wordBreakDp(s, wordMap, 0, memo)
}

func wordBreakDp(s string, wordMap map[string]bool, idx int, memo []int) bool {
	if len(s) == idx {
		return true
	}

	if memo[idx] != -1 {
		if memo[idx] == 0 {
			//0 ~ idx 无法凑出
			return false
		} else {
			//0 ~ idx 可以凑出
			return true
		}
	}

	for i := idx; i < len(s); i++ {
		ss := s[idx : i+1]
		if wordMap[ss] {
			if wordBreakDp(s, wordMap, i+1, memo) {
				memo[idx] = 1
				return true
			}
		}
	}
	memo[idx] = 0
	return false
}

// 140. 单词拆分 II
func wordBreakII(s string, wordDict []string) []string {
	wordMap := make(map[string]bool)
	for _, word := range wordDict {
		wordMap[word] = true
	}
	words := make([]string, 0)
	sentences := make([]string, 0)
	dpWordBreakII(s, wordMap, 0, words, &sentences)
	return sentences
}

func dpWordBreakII(s string, wordMap map[string]bool, idx int, words []string, sentences *[]string) {
	if idx == len(s) {
		sentence := strings.Join(words, " ")
		*sentences = append(*sentences, sentence)
		return
	}
	for i := idx; i < len(s); i++ {
		ss := s[idx : i+1]
		if wordMap[ss] {
			words = append(words, ss)
			dpWordBreakII(s, wordMap, i+1, words, sentences)
			words = words[:len(words)-1]
		}
	}
}

// 115. 不同的子序列 https://blog.csdn.net/fdl123456/article/details/124938272
func numDistinct(s string, t string) int {
	mem := make([][]int, len(s))
	for i, _ := range mem {
		mem[i] = make([]int, len(t))
		for j, _ := range mem[i] {
			mem[i][j] = -1
		}
	}
	return dpND(s, 0, t, 0, mem)
}

func dpND(s string, i int, t string, j int, mem [][]int) int {
	if j == len(t) {
		//只有s[i] == t[j]时, j才会往后走(j+1), 所以j到达len(t), 说明匹配到了
		return 1
	}
	//说明s[i..]已经不能凑出t[j..]
	if len(s)-i < len(t)-j {
		return 0
	}
	if mem[i][j] != -1 {
		return mem[i][j]
	}
	if s[i] == t[j] {
		//当s[i] == t[j]时, 可以选择i j都匹配， 也可以i不匹配
		mem[i][j] = dpND(s, i+1, t, j, mem) + dpND(s, i+1, t, j+1, mem)
	} else {
		//当s[i] != t[j]时, 只能i不匹配
		mem[i][j] = dpND(s, i+1, t, j, mem)
	}
	return mem[i][j]
}

// 迭代, 自下而上, 从base case推出所有结果
func numDistinctV2(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	//dp[i][j]表示 s[i:]的子序列中包含t[j:]的数量
	dp := make([][]int, m+1)
	//dp[i][n]表示s[i:]的子序列中包含t[n:]（空字符串）的数量, 肯定包含, 都为1
	//dp[m][j]表示s[m:]（空字符串）的子序列中包含t[j:]的数量, 肯定没有, 都为0
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
		dp[i][n] = 1
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				//如果s[i] == t[j], 有两种情况. s[i:]的子序列包含t[j:]的数量 + s[i+1:]的子序列包含t[j:]的数量
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				//如果s[i] != t[j], 只有一种情况. s[i+1:]的子序列包含t[j:]的数量
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return dp[0][0]
}

// 72. 编辑距离
/*
思路： s[i]==t[j], 不需要操作,由s[i-1],t[j-1]转化而来, dp[i-1][j-1];
      s[i]!=t[j], 由s增, s删, s改三种情况的最小值转换操作步骤+1，即min{dp[i][j-1], dp[i-1][j], dp[i-1][j-1]+1}

base case：
i == 0, 只需要给s新增元素, dp[0][j] = j
j == 0, 只需要给s删除元素, dp[i][0] = i
*/
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	for i := 0; i <= len(word1); i++ {
		for j := 0; j <= len(word2); j++ {
			if i == 0 {
				dp[i][j] = j
				continue
			}
			if j == 0 {
				dp[i][j] = i
				continue
			}
			// 0 这个位置只计数
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = Min(Min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

/*
最长公共子序列问题
https://mp.weixin.qq.com/s/ZhPEchewfc03xWv9VP3msg

1143. 最长公共子序列
思路:
s[i] == t[j], 说明dp[i][j] = dp[i-1][j-1] + 1
s[i] != t[j], 两种情况, 要么s退后一步, 要么t退后一步, dp[i][j] = min{dp[i-1][j], dp[i][j-1]}

base case: 当i == 0时, dp[0][j] = 0; 当j == 0时, dp[i][0] == 0

583. 两个字符串的删除操作
712. 两个字符串的最小ASCII删除和

备注: 这三题就是"编辑距离"的简化版
*/
//712. 两个字符串的最小ASCII删除和
func TestMDS(t *testing.T) {
	minimumDeleteSum("abc", "abc")
}

func minimumDeleteSum(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}
	for i := 0; i <= len(s1); i++ {
		for j := 0; j <= len(s2); j++ {
			if i == 0 {
				//注意这里累加到j
				for m := 0; m < j; m++ {
					dp[i][j] += int(s2[m])
				}
				fmt.Printf("dp[%d][%d] = %d\n", i, j, dp[i][j])
				continue
			}
			if j == 0 {
				//注意这里累加到i
				for m := 0; m < i; m++ {
					dp[i][j] += int(s1[m])
				}
				fmt.Printf("dp[%d][%d] = %d\n", i, j, dp[i][j])
				continue
			}
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
				fmt.Printf("dp[%d][%d] = %d\n", i, j, dp[i][j])
			} else {
				dp[i][j] = Min(dp[i-1][j]+int(s1[i-1]), dp[i][j-1]+int(s2[j-1]))
				fmt.Printf("dp[%d][%d] = %d\n", i, j, dp[i][j])
			}
		}
	}
	return dp[len(s1)][len(s2)]
}

// 516. 最长回文子序列
// https://mp.weixin.qq.com/s/zNai1pzXHeB2tQE6AdOXTA
/*
思路: 设dp[i][j]为s[i...j]的最长回文子序列,
如果s[i]==s[j], s[i]和s[j]都可以加入到回文中, dp[i][j] = d[i+1][j-1]+2,
如果s[i]!=s[j], 两种情况, 要么s[i]要么s[j], dp[i][j] = max{d[i+1][j], dp[i][j-1]}

base case: i == j, 说明只有一个字符, dp[i][j] == 1,i > j, 不可能, dp[i][j] == 0


1312. 让字符串成为回文串的最少插入次数 与本题类似
*/
func longestPalindromeSubseq(s string) int {
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := 0; j < len(s); j++ {
			if i == j {
				//如果只有一个字符
				dp[i][j] = 1
				continue
			}
			if i > j {
				//不存在的位置
				dp[i][j] = 0
				continue
			}

			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = Max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}

/*
0-1背包问题

416. 分割等和子集
322. 零钱兑换 也可以用二维解法。一维数组更简单(完全背包问题)

二者不同: 分割子集数组元素只能用一次, 零钱兑换可以用无数次
*/
func canPartition(nums []int) bool {
	// 求所有元素和的一半sum, 判断nums中是否有子集能组成sum。转化为背包问题
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	sum = sum / 2
	// dp[i][j] 表示前i个元素能否凑成j
	dp := make([][]bool, len(nums)+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}

	for i := 0; i <= len(nums); i++ {
		for j := 0; j <= sum; j++ {
			if i == 0 {
				//没有元素肯定装不满
				dp[i][j] = false
				continue
			}
			if j == 0 {
				//相当于j是装满的
				dp[i][j] = true
				continue
			}
			if j-nums[i-1] < 0 {
				// 没有空间选第i个元素, 继承前i-1个元素的值
				dp[i][j] = dp[i-1][j]
			} else {
				// 有空间, 选择nums[i-1]和不选nums[i-1]做为子集。
				// 选   dp[i-1][j]
				// 不选 dp[i-1][j-nums[i-1]]
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[len(nums)][sum]
}

func coinChangeV2(coins []int, amount int) int {
	//dp[i][j]表示前i个元素凑成j所需的最少硬币
	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}
	for i := 0; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if i == 0 {
				dp[i][j] = amount + 1
				continue
			}
			if j == 0 {
				dp[i][j] = 0
				continue
			}

			if j-coins[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				//todo 思考为啥这里是dp[i] 上一题是dp[i-1]
				//不使用  dp[i-1][j]
				//使用    dp[i][j-coins[i-1]] + 1
				dp[i][j] = Min(dp[i-1][j], dp[i][j-coins[i-1]]+1)
			}
		}
	}
	if dp[len(coins)][amount] == amount+1 {
		return -1
	}
	return dp[len(coins)][amount]
}
