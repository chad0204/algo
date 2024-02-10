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
	fmt.Println(coinChangeDp([]int{2, 5, 10, 1}, 27))
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

func coinChangeDp(coins []int, amount int) int {
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
	/*
	   dp[i][j][0] 表示第i天, 交易j笔, 未持有的最大利润
	   dp[i][j][1] 表示第i天, 交易j笔, 持有股票时的最大利润

	   前一天未持有, 今天无操作; 前一天持有, 今天卖了一笔, 笔数不变, 笔数在买的时候扣掉
	   dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1] + prices[i-1])
	   前一天持有, 今天无操作; 前一天未持有, 今天买了一笔, 说明前一天的笔数是j-1(今天是j)
	   dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0] - prices[i-1])


	   dp[0][j][0] = 0
	   dp[0][j][1] = math.MinInt32

	*/
	n := len(prices)
	dp := make([][][]int, n+1)
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= k; j++ {
			if i == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = math.MinInt32
				continue
			}
			if j == 0 {
				dp[i][j][0] = 0
				dp[i][j][1] = math.MinInt32
				continue
			}
			dp[i][j][0] = Max(dp[i-1][j][0], dp[i-1][j][1]+prices[i-1])
			dp[i][j][1] = Max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i-1])
		}
	}
	return dp[n][k][0]
}

// k = 1 “今天买了不能卖” 和 k无限次的区别就是一天不能同时买卖
// 121. 买卖股票的最佳时机
func maxProfit1(prices []int) int {

	/*
			1次交易
			买 卖 无操作
			持有 未持有

		   第i天不持有股票
		   dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
		   第i天持有股票, 之前就持有今天无操作; 要么之前不持有今天买的, 由于只能买一次, 那么之前利润一定0, 今天0 - price[i]
		   dp[i][1] = max(dp[i-1][1], -prices[i])


		   i == 0
		   dp[0][0] = 0
		   dp[0][1] = 不可能
	*/
	dp := make([][]int, len(prices)+1)
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	for i := 0; i <= len(prices); i++ {
		if i == 0 {
			dp[0][0] = 0
			dp[0][1] = math.MinInt32
			continue
		}
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i-1])
		dp[i][1] = Max(dp[i-1][1], -prices[i-1])
	}
	return dp[len(prices)][0]
}

// k = +∞
// 122. 买卖股票的最佳时机 II
func maxProfitInfinity(prices []int) int {
	/**

	  dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i-1])
	  dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i-1])

	  dp[0][0] = 0
	  dp[0][1] = -math.MinInt32

	*/
	dp := make([][]int, len(prices)+1)
	for i := range dp {
		dp[i] = make([]int, 2)
	}

	for i := 0; i <= len(prices); i++ {
		if i == 0 {
			dp[0][0] = 0
			dp[0][1] = math.MinInt32
			continue
		}
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i-1])
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i-1])
	}
	return dp[len(prices)][0]
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
思路：在i位置上lis等于任意比nums[i]小的元素的lis+1, 取最大值。dp[i] = max{dp[j]+1}, 0<=j<i,nums[j]<nums[i]
base case dp[0] = 1, 且任意位置的最小值都是1
*/
func lengthOfLIS(nums []int) int {
	//dp[i]表示以nums[i]为结尾的最长子序列的长度
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

func TestLengthOfLISV2(t *testing.T) {
	lengthOfLISV2([]int{10, 9, 2, 5, 3, 7, 21, 18})
	//lengthOfLISV2([]int{7, 8, 9, 1, 2, 3, 5, 4})
}

/**

解法一: 朴素dp

dp[i]表示以nums[i]为结尾的最长递增子序列的长度
初始化dp[i] = 1

每个dp[i] = max(dp[j]+1, dp[i]), nums[j] < nums[i], k属于[0,i)



解法二：二分查找+贪心

dp[i]表示以长度为i+1的最长递增子序列的最小值）
最小值: 比如1254中长度为3的递增子序列为 1,2,5; 1,2,4。长度为3递增子序列的最小值应该是4而不是5
i+1, 因为二分查找确定位置时, 得到的结果是当前长度的最后一位i+1, 比如124,6, 迭代到6的时候二分查找[0,3]插入位置是3, 等于上一个长度len, 所以i+1更方便。


初始化dp[i] = 0

1 2 5 6
已知的长度为3的值为5, dp[2] = 5,
如果当前迭代到的数组值是6, 那么二分查找6在dp中的位置, 也就是找[0, 2+1]中的位置, 结果是3, dp[3] = 6, dp[3] = 6, 插入的位置正好是3, 说明当前元素比之前的元素都大, 所以新长度为3+1

1 2 5 3
如果当前迭代到的数组值是3, 那么二分查找3在dp中的位置, 也就是找[0, 2+1]中的位置, 结果是2, dp[2] = 3。 相当于3在位置2取代了5(肯定取小的), 插入的位置是2, 长度不变

demo1:
比如 10, 9, 2, 5, 3, 7, 21, 18
dp[0] = 10, len = 1
dp[0] = 9, len = 1
dp[0] = 2, len = 1
得到dp[0] = 2

dp[1] = 5, len = 2
dp[1] = 3, len = 2
2,5
2,3
得到dp[1] = 3

dp[2] = 7, len = 3
2,3,7
得到dp[2] = 7


dp[3] = 21, len = 4
dp[3] = 18, len = 4
2,3,7,21
2,3,7,18
得到dp[3] = 18

demo2:
nums = {7, 8, 9, 1, 2, 3, 5, 4}, len=0

nums=7, index = 0, dp[0] = 7, len = 1, 规则1 二分查找得到的index等于上一个长度0, 表示nums[i]比之前的元素都大, 长度+1
nums=8, index = 1, dp[1] = 8, len = 2, 规则1
nums=9, index = 2, dp[2] = 9, len = 3, 规则1
nums=1, index = 0, dp[0] = 1, len = 3, 规则2二分查找得到的index<len, 表示nums[i]比[0,index]的元素都小, 长度不变, 替换原来index的元素。这里可以理解为长度为1的结尾最小值应该是1而不是之前的7
nums=2, index = 1, dp[1] = 2, len = 3, 规则2
nums=3, index = 2, dp[2] = 3, len = 3, 规则2
nums=5, index = 3, dp[3] = 5, len = 4, 规则1
nums=4, index = 3, dp[3] = 4, len = 4, 规则2



gpt解释：
1. 我们维护一个数组 dp，其中 dp[i] 表示长度为 i+1 的递增子序列的末尾元素的最小值。注意，dp 数组并不一定是一个有效的递增序列，但是 dp 的长度就是当前最长递增子序列的长度。
2. 我们遍历输入数组 nums，对于每个元素 num，我们使用二分查找找到 dp 中第一个大于或等于 num 的位置 index。如果 index 等于当前最长递增子序列的长度，说明 num 大于当前递增子序列的最大值，因此我们更新当前递增子序列的长度。
3. 如果 index 小于当前最长递增子序列的长度，说明我们可以用 num 替换掉 dp[index]，因为 num 较小，有可能构成更长的递增子序列。
4. 最终，dp 数组的长度即为最长递增子序列的长度。


35. 搜索插入位置

*/

// 300. 最长递增子序列（动态规划 + 二分查找，清晰图解）
func lengthOfLISV2(nums []int) int {
	n := len(nums)
	//maxLenNums[i]表示长度为i+1的递增子序列的最后一位的值
	maxLenNums := make([]int, n)
	//已有的最长序列
	length := 0 //可以省略, 直接用len(maxLenNums)
	//寻找nums[i]的插入位置
	for i := 0; i < n; i++ {
		// 使用二分查找, 找到nums[i]在dp中的插入位置
		idx := searchInsertIdx(maxLenNums, nums[i], length)
		if idx == length {
			//说明插入位置是已有最长序列的尾部, 即nums[i]大于当前已记录的递增序列的值的最大值，更新递增序列长度。
			length++
		}
		//替换长度为length的位置的
		maxLenNums[idx] = nums[i]
	}
	return length
}

func searchInsertIdx(dp []int, target int, length int) int {
	left, right := 0, length-1
	for left <= right {
		mid := left + (right-left)/2
		if target > dp[mid] {
			left = mid + 1
		} else if target < dp[mid] {
			right = mid - 1
		} else {
			left = mid
			break
		}
	}
	return left
}

// 931. 下降路径最小和. 思考： 是不是可以给长和宽都多一列(值为最大值, 这样就可以避免边界溢出)。试试自上而下的递归？
func minFallingPathSum(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i := range dp {
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

func TestWord(t *testing.T) {

	wordBreak("leetcode", []string{"leet", "code"})
}

// 139. 单词拆分
func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, v := range wordDict {
		wordMap[v] = true
	}
	/**
	  dp[i] 表示字符串s的前i个字符组成的字符串s[0..i−1]是否能被空格拆分成若干个字典中出现的单词
	  dp[i] = dp[j] && word contains dp[j:i] j属于[0,i)
	  dp[0] = true, 用来表示空字符串。比如s = leet, word = {"leet", "code"}, s可以被拆分成"" "leet"

	  如果用dp[0]表示第一个字符, 代码比较难写

	*/
	n := len(s)
	dp := make([]bool, n+1)
	//空字符串为true
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
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
	for i := range mem {
		mem[i] = make([]int, len(t))
		for j := range mem[i] {
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
	for i := range dp {
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

// 5. 最长回文子串
func longestPalindromeDP(s string) string {
	/**
	  dp[i][j]表示s[i:j]是否是回文子串
	*/
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	maxLen := 1
	start := 0
	end := 0
	for r := 0; r < n; r++ {
		for l := 0; l <= r; l++ {
			//base case 同一个字符, 肯定是回文
			if l == r {
				dp[l][r] = true
				continue
			}
			//字符相等
			if s[r] == s[l] {
				//相邻, 是回文
				if r-l == 1 {
					dp[l][r] = true
				}
				//前一个状态也是回文
				if dp[l+1][r-1] {
					dp[l][r] = true
				}
			}
			//根据最大长度更新位置
			if dp[l][r] && r-l+1 > maxLen {
				maxLen = r - l + 1
				start = l
				end = r
			}
		}
	}
	return s[start : end+1]
}

/*
0-1背包问题

416. 分割等和子集  1 5 11 5    11

1 5 1 5

322. 零钱兑换 也可以用二维解法。一维数组更简单(完全背包问题)

二者不同: 分割子集(01背包)数组元素只能用一次, 零钱兑换(完全背包)可以用无数次

假设背包的容量为5。有四个物品，它们的重量和价值分别为:
物品1: 重量 w1 = 2, 价值 v1 = 3
物品2: 重量 w2 = 1, 价值 v2 = 2
物品3: 重量 w3 = 3, 价值 v3 = 4
物品4: 重量 w4 = 2, 价值 v4 = 2

0-1背包: 考虑第i个元素时, 加上v[i]之后求j-w[i]时, 不能再带上i
dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i]] + v[i])
完全背包: 考虑第i个元素时, 加上v[i]之后求j-w[i]时, 可以再带上i, 能用多次
dp[i][j] = max(dp[i-1][j], dp[i][j-w[i]] + v[i])
*/
func TestCP(t *testing.T) {
	canPartition([]int{1, 2, 3})
}

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
			if j == 0 {
				//相当于j是装满的
				dp[i][j] = true
				continue
			}
			if i == 0 {
				//没有元素肯定装不满
				dp[i][j] = false
				continue
			}
			// 可以理解为i是从1计数的
			if j-nums[i-1] < 0 {
				// 不能把nums[i]算入子集, 继承前i-1个元素的值
				dp[i][j] = dp[i-1][j]
			} else {
				// 有空间, 选择nums[i-1]和不选nums[i-1]做为子集。
				// 不把nums[i-1]算入子集   dp[i-1][j]
				// 把nums[i-1]算入子集 dp[i-1][j-nums[i-1]], 就要看前i-1个元素能不能凑成j-nums[i-1]
				/*
						思考为啥这里是dp[i-1][j-nums[i-1]] 而不是dp[i][j-nums[i-1]]。

					dp[i-1][j-nums[i-1]] 表示算上当前元素, 能不能凑成j。那么就看前i-1个元素能不能凑成j-nums[i-1]。如果用dp[i][j-nums[i-1]], 那么nums[i-1]就重复计算了, 相当于算上nums[i-1]去计算j-nums[i-1]
				*/
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
		for j := 0; j <= amount; j++ {
			if j == 0 {
				dp[i][j] = 0
				continue
			}
			if i == 0 {
				dp[i][j] = amount + 1
				continue
			}

			if j-coins[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				//不使用  dp[i-1][j]
				//使用    dp[i][j-coins[i-1]] + 1 使用则要把i加进来, 然后硬币数+1。而01背包的nums[i]算入子集, 则不能把i加进来算j-nums[i-1]
				dp[i][j] = Min(dp[i-1][j], dp[i][j-coins[i-1]]+1)
			}
		}
	}
	if dp[len(coins)][amount] == amount+1 {
		return -1
	}
	return dp[len(coins)][amount]
}

// 518. 零钱兑换 II
func change(amount int, coins []int) int {
	//dp[i][j] 表示前i个元素, 能凑成j的方式的个数
	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}
	for i := 0; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if j == 0 {
				dp[i][j] = 1
				continue
			}
			if i == 0 {
				dp[i][j] = 0
				continue
			}
			if j-coins[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}
	return dp[len(coins)][amount]
}

// 494. 目标和
// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247485700&idx=1&sn=433fc5ec5e03a86064d458320332a688&chksm=9bd7f70caca07e1aad658333ac05df501796862a418d8f856b12bb6ca73a924552901ec86d9b&cur_album_id=1318881141113536512&scene=189#wechat_redirect
func TestFS(t *testing.T) {
	findTargetSumWays([]int{0, 0, 1}, 1)
}

func findTargetSumWays(nums []int, target int) int {
	/*
	   赋+的子集为A, 赋-的子集为B
	   sum(A) - sum(B) = target
	   sum(A) = target + sum(B)
	   2*sum(A) = target + sum(B) + sum(A)
	   sum(A) = (target + sum(nums)) / 2
	   存在多少子集A, 使得A正好能凑够(target + sum(nums)) / 2
	*/
	sum := 0
	for _, v := range nums {
		sum += v
	}
	//防止target是负数
	sum = abs(target) + sum
	if sum%2 == 1 || sum < abs(target) {
		return 0
	}
	sum = sum / 2
	//dp[i][j]表示前i个元素，能凑成j的方法的数量
	dp := make([][]int, len(nums)+1)
	for i := range dp {
		dp[i] = make([]int, sum+1)
	}

	//先初始化,j == 0
	for i := 0; i <= len(nums); i++ {
		dp[i][0] = 1
	}
	for i := 1; i <= len(nums); i++ {
		//依然j==0依然要算, 防止nums[0]==0的情况, +0和-0是两种情况
		for j := 0; j <= sum; j++ {
			if j-nums[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[len(nums)][sum]
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// 64. 最小路径和
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m := len(grid)    //3
	n := len(grid[0]) //3
	//dp[i][j]表示 i,j时最小
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 {
				dp[i][j] = dp[i][j-1] + grid[i-1][j-1]
				continue
			}
			if j == 1 {
				dp[i][j] = dp[i-1][j] + grid[i-1][j-1]
				continue
			}
			dp[i][j] = Min(dp[i-1][j], dp[i][j-1]) + grid[i-1][j-1]
		}
	}
	return dp[m][n]
}

/*
42. 接雨水

1. i往两边延伸, 遇到两边的最大高度l, r, 最大高度max = min(l, r)
2. i能接的水量 max - height[i]

0,1,0,2,1,0,1,3,2,1,2,1

4,2,0,3,2,5

暴力解法时间复杂度 n*n

*/

func trap(height []int) int {
	n := len(height)
	temp := make([]int, n)
	for i := 1; i < n-1; i++ {
		lMax := 0
		for j := i - 1; j >= 0; j-- {
			if height[j] > height[i] {
				lMax = Max(height[j], lMax)
			}
		}
		rMax := 0
		for j := i + 1; j < n; j++ {
			if height[j] > height[i] {
				rMax = Max(height[j], rMax)
			}
		}
		if Min(lMax, rMax) > height[i] {
			temp[i] = Min(lMax, rMax) - height[i]
		}
	}
	res := 0
	for _, v := range temp {
		res += v
	}
	return res
}

/*
1. i往两边延伸, 遇到两边的最大高度l, r, 最大高度max = min(l, r)
2. i能接的水量 max - height[i]

0,1,0,2,1,0,1,3,2,1,2,1

4,2,0,3,2,5

dp思路: 先分别计算i的左右最大位置
*/
func trapDp(height []int) int {
	n := len(height)
	//leftMax[i]表示i以及i左边的所有位置中的height最大值
	leftMax := make([]int, n)

	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = Max(leftMax[i-1], height[i])
	}

	//rightMax[i]表示i以及i右边的所有位置中的height最大值
	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = Max(rightMax[i+1], height[i])
	}

	res := 0
	for i := 0; i < n; i++ {
		//if Min(leftMax[i], rightMax[i]) > height[i] {
		res += Min(leftMax[i], rightMax[i]) - height[i]
		//}
	}
	return res
}
func TestMaxSubArray(t *testing.T) {
	maxSubArray([]int{-1, -2})
}

// 53. 最大子数组和
func maxSubArray(nums []int) int {
	n := len(nums)
	//dp[i]表示以nums[i]为结尾的连续数组的最大值
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i] = nums[i]
			continue
		}
		if dp[i-1] > 0 {
			//dp[i-1]是正值, 无论num[i]是正是负都加上,
			//注意这里会让值变小, 没事, 因为以i结尾的连续值就是会小, 后面会遍历一遍取最大值
			dp[i] = dp[i-1] + nums[i]
		} else if nums[i] < dp[i-1] {
			//dp[i-1]是负数, 当前值比dp[i-1]小, 从i重新开始
			dp[i] = nums[i]
		} else {
			//dp[i-1]是负数, 当前值比dp[i-1]大, 从i重新开始
			dp[i] = nums[i]
		}
	}

	res := math.MinInt32
	for _, v := range dp {
		res = Max(res, v)
	}
	return res
}
