package datastruct

import (
	"testing"
)

/*
有一个背包, 容量为weight, 有一堆物品, 重量为w[i], 价值为v[i]

01背包(物品只能使用一次)
416. 分割等和子集
背包容量为target, 物品重量为nums[i], 物品价值为nums[i], 求物品放满后的最大价值能不能到target

完全背包问题：
322. 零钱兑换
背包容量为amount, 物品重量为nums[i], 物品价值为1（一枚硬币单位是一个）, 求物品放满背包的最小的价值

518. 零钱兑换 II
背包容量为amount, 物品重量为nums[i], 与价值无关, 求物品放满背包的组合数


0-1背包: 考虑第i个物品时, 选择第i个物品后之后求j-w[i]时, 不能再带上第i个物品
dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i]] + v[i])
完全背包: 考虑第i个物品时, 选择第i个物品之后求j-w[i]时, 可以再带上i, 能用多次
dp[i][j] = max(dp[i-1][j], dp[i][j-w[i]] + v[i])


*/

/*
416. 分割等和子集
*/
func TestCP(t *testing.T) {
	canPartition([]int{1, 2, 3})
}
func canPartition(nums []int) bool {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	/**
	  判断nums能不能凑出target
	  dp[i][j]表示前i个元素凑出j得到的最大值
	  把nums[i]放进来， dp[i][j] = dp[i-1][j-nums[i-1]] + nums[i-1]
	  不把nums[i]放进， dp[i][j] = dp[i-1][j]

	  动态转移方程:
		dp[i][j] = Max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i]), i从0开始遍历, i-1表示物品位置, i == 0表示没有物品

		1. dp[i-1][j], 表示不选择nums[i-1], 结果就是前i-1个物品凑出的最大价值
		2. dp[i-1][j-nums[i]]+nums[i], 表示选择nums[i-1], 结果就是前i个物品凑出的最大价值, 由于选择了nums[i-1], 那么之前凑出j-nums[i]就不能选择nums[i-1]
		3. 如果j < nums[i-1], 放不了, 只能选择不选择

	  base case:
		i == 0, 表示没有物品, 价值为0
		j == 0, 表示背包容量为0, 价值为0

	*/
	n := len(nums)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= target; j++ {
			if i == 0 {
				dp[i][j] = 0
				continue
			}
			if j == 0 {
				dp[i][j] = 0
				continue
			}
			if j < nums[i-1] {
				//想放也放不了
				dp[i][j] = dp[i-1][j]
			} else {
				//这里取选择和不选的最大值
				dp[i][j] = Max(dp[i-1][j] /*不选nums[i-1]*/, dp[i-1][j-nums[i-1]]+nums[i-1] /*选择nums[i-1]*/)
			}
		}
	}
	return dp[n-1][target] == target
}

/*
322.零钱兑换
*/
func coinChange(coins []int, amount int) int {
	/**
	  判断coins凑出amount的最小价值
	  dp[i][j]表示前i个元素凑出j得到的最小值

	  动态转移方程:
		dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i-1]]+1), i从0开始遍历, i-1表示物品位置, i == 0表示没有物品

		1. dp[i-1][j], 表示不选择coins[i-1], 结果就是前i-1个物品凑出的最大价值
		2. dp[i][j-coins[i-1]]+1, 表示选择coins[i-1], 结果就是前i个物品凑出的最大价值, 但是由于可以选无数次, 即使选择了nums[i-1], 凑出j-coins[i-1]也可以选择coins[i-1]
		3. 如果j < nums[i-1], 放不了, 只能选择不选择

	  base case:
		i == 0, 表示没有物品, 价值为0
		j == 0, 表示背包容量为0, 价值为amount+1, 设置个最大值
	*/
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
func changeII(amount int, coins []int) int {
	/**
	  判断coins凑出amount的方式数
	  dp[i][j]表示前i个元素凑出j的方式数

	  动态转移方程:
		dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]], i从0开始遍历, i-1表示物品位置, i==0表示没有物品
		组合数就是选择和不选择两种情况之和

		1. dp[i-1][j], 表示不选择coins[i-1], 结果是前i-1个物品凑出j的方式数
		2. dp[i][j-coins[i-1]], 表示选择coins[i-1], 结果是算上i之后凑出j的方式数, 但是由于可以选无数次, 即使选择了coins[i-1], 凑出j-coins[i-1]也可以选择coins[i-1]
		3. 如果j < coins[i-1], 放不了, 只能选择不选择

	  base case:
		i == 0, 表示没有物品, 无法凑出
		j == 0, 表示背包容量为0, 方式只有一种, 就是啥也不装
	*/
	dp := make([][]int, len(coins)+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
	}
	for i := 0; i <= len(coins); i++ {
		for j := 0; j <= amount; j++ {
			if i == 0 {
				dp[i][j] = 0
				continue
			}
			if j == 0 {
				dp[i][j] = 1
				continue
			}
			if j-coins[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				//选和不选两种情况之和
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}
	return dp[len(coins)][amount]
}

// 下面两种是一维数组解法
func coinChangeOne(coins []int, amount int) int {
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

func changeOne(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}
	return dp[amount]
}
