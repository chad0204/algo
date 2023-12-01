package datastruct

import (
	"fmt"
	"testing"
)

/*
给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。

示例 1：
输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
示例 2：
输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。


 0 1 2 3 4
[3,2,1,0,4]

i = 0; l = 3
i = 1; l = 3
i = 2; l = 3
i = 3; l = 3
i = 4; l = 8


最远能条多远，如果最远比数组长度长，表示能跳到最后

*/
func TestA(t *testing.T) {
	fmt.Println(canJumpDp2([]int{3, 2, 1, 0, 4}))
}

func canJump(nums []int) bool {
	mem := make([]int, len(nums))
	for i := range mem {
		mem[i] = -1
	}
	return canJumpHelper(nums, 0, mem) >= len(nums)-1
}

func canJumpHelper(nums []int, idx int, mem []int) int {
	if idx >= len(nums)-1 {
		return idx
	}
	if mem[idx] != -1 {
		return mem[idx]
	}
	res := -1
	for i := 1; i <= nums[idx]; i++ {
		res = Max(res, canJumpHelper(nums, idx+i, mem))
	}
	mem[idx] = res
	return res
}

//一维数组没搞出来
func canJumpDp2(nums []int) bool {
	//dp[i]表示前i个元素, 能到达的最大距离
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			dp[i] = nums[i]
			continue
		}
		dp[i] = Max(dp[i-1], i+nums[i])
		//if dp[i] <= i {
		//	return false
		//}
	}
	return dp[len(nums)-1] >= len(nums)-1
}

func canJumpGreedy(nums []int) bool {
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		if maxLen < i {
			return false
		}
		maxLen = Max(maxLen, i+nums[i])
	}
	return maxLen >= len(nums)-1
}
