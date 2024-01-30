package datastruct

import "testing"

func TestCoint(t *testing.T) {
	coinChangeIA([]int{1, 2, 5}, 11)
}

func coinChangeIA(coins []int, amount int) int {
	/**
	  dp[i]表示凑够i最少需要多少硬币

	  dp[0] = 0, 其他初始化成amount+1

	  dp[i] = min(dp[i-c] + 1, dp[i]); i>=c

	*/
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for _, c := range coins {
		//用前i个硬币凑出amount的最小值
		for j := c; j <= amount; j++ {
			if dp[j-c] == amount+1 {
				continue
			}
			dp[j] = Min(dp[j-c]+1, dp[j])
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func coinChangeIB(coins []int, amount int) int {
	/**
	  dp[i]表示凑够i最少需要多少硬币

	  dp[0] = 0, 其他初始化成amount+1

	  dp[i] = min(dp[i-c] + 1, dp[i]); i>=c

	*/
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for _, c := range coins {
		for j := c; j <= amount; j++ {
			if dp[j-c] == amount+1 {
				continue
			}
			dp[j] = Min(dp[j-c]+1, dp[j])
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func changeII(amount int, coins []int) int {
	//dp[j]表示使用硬币凑成i的组合数
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		//使用当前nums[i]凑成所有j的数量
		for j := 1; j <= amount; j++ {
			//上面从coins[i]开始遍历, 就不用判断了
			if j >= coins[i] {
				dp[j] += dp[j-coins[i]]
			}
		}
	}
	return dp[amount]
}
