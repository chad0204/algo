package datastruct

import (
	"fmt"
	"sort"
	"testing"
)

/*
贪心算法
比如零钱兑换, 先用最大的硬币凑, 就是一种贪心选择思路, 只考虑局部最优解, 但是忽略了子问题直接的关系
10 9 4 1 凑 13
先用最大的硬币10, 那么结果就是 10 1 1 1 = 4, 但是先取10, 会导致无法使用4和9。所以最终可能会得不到最优解。

所以是否能用贪心算法, 要看问题是否具有贪心选择性质, 也就是通过局部最优解获得全局最优解。

https://mp.weixin.qq.com/mp/appmsgalbum?action=getalbum&__biz=MzAxODQxMDM0Mw==&scene=1&album_id=2165269406745001984&count=3#wechat_redirect


区间的三种形式：
部分重叠
全部包含
无重叠

*/

// 435. 无重叠区间
/*
动态规划思路：
移除最少区间剩下的不重叠 == 区间数 - 最多不重叠区间数
求： 最多不重叠区间数

按照区间左点排序(右也可)

dp[i]表示，前i个区间中, 如果有区间j右区间比i的左区间小, 那么区间j满足不重叠, dp[i]为这些j区间中最大的结果+1
dp[i] = max{dp[j]} + 1 (j < i, interval[j][1] <= interval[i][0])
res = len(intervals) - max(dp)

*/
func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	n := len(intervals)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		//至少一个
		dp[i] = 1
		for j := 0; j < i; j++ {
			if intervals[i][0] >= intervals[j][1] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
	}
	res := 0
	for _, v := range dp {
		res = Max(v, res)
	}
	return n - res
}

/**
我们希望每次选择的区间的结束时间越早越好，这是因为结束得越早，剩下的空间就越大，
从而可以容纳更多的区间。这种贪心选择性质使得贪心算法在这个问题上可以有效地找到最优解。


比如零钱兑换, 先用最大的硬币凑, 就是一种贪心思路, 只考虑局部最优解, 但是忽略了子问题直接的关系
10 9 4 1 凑 13
先用最大的硬币10, 那么结果就是 10 1 1 1 = 4, 但是先取10, 会导致无法使用4和9。所以最终可能会得不到最优解。

所以是否能用贪心算法, 要看问题是否具有贪心选择性质。

在无重叠区间的问题中，贪心选择性质指的是在每一步选择中，都选择局部最优的区间，希望通过这样的选择，最终得到全局最优的解。

具体来说，对于这个问题，我们希望尽量保留更多的区间，即尽量选择那些结束时间较早的区间。这是因为结束得越早，剩余的空间就越大，能够容纳更多的区间。

假设有一组区间 [a1, b1], [a2, b2], ..., [an, bn]，并且这些区间已经按照结束时间的升序排列。在贪心选择中，我们首先选择结束时间最早的区间 [a1, b1]。这是贪心选择的一部分，因为我们选择了当前状态下的局部最优解。

然后，在剩余的区间中，我们继续选择下一个结束时间最早的区间。每一步都选择当前状态下结束时间最早的区间，就是贪心选择。

贪心选择性质在这里成立的原因是，选择了结束时间最早的区间之后，剩下的空间更大，因此在后续的选择中有更多的可能性容纳更多的区间。这种选择方式使得贪心算法在这个问题上能够有效地找到最优解，即保留尽量多的区间。


*/
// 贪心算法, 按照start排序
func eraseOverlapIntervalsStart(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	x_end := intervals[0][1]
	count := 1
	for _, interval := range intervals {
		if interval[0] >= x_end {
			count++
			x_end = interval[1]
		} else {
			x_end = Min(x_end, interval[1])
		}
	}
	return len(intervals) - count
}

// 贪心算法, 按照end排序
func eraseOverlapIntervalsEnd(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	x_end := intervals[0][1]
	count := 1
	for _, interval := range intervals {
		//下一个区间的start比之前的区间的end大, 计数加1并更新end
		if interval[0] >= x_end {
			count++
			x_end = interval[1]
		}
	}
	return len(intervals) - count
}

// 452. 用最少数量的箭引爆气球
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	end := points[0][1]
	count := 1
	for _, point := range points {
		if point[0] > end {
			count++
			end = point[1]
		}
	}
	return count
}

// 1024. 视频拼接

/*
部分重叠
全部包含
无重叠

先按照起点排序, 如果起点相同, 那么选择最长的一定最优, 也就是起点相同选择终点最大的。
当前选择x[s, e], 那么下一个被选中的一定起点小于s, 终点最大的区间

-----------------
------------

		--------------------
		 -----------------
	                     -------
*/
func videoStitching(clips [][]int, time int) int {
	sort.Slice(clips, func(i, j int) bool {
		return clips[i][0] < clips[j][0]
	})
	end := 0
	count := 0
	maxEnd := 0
	for i := 0; i < len(clips); {
		//有间隙, 直接失败
		if clips[i][0] > end {
			return -1
		}
		maxEnd = clips[i][1]
		i++
		//遍历i之后, 所有start比i的end小的区间, 找出end最大的
		for i < len(clips) && clips[i][0] <= end {
			maxEnd = Max(maxEnd, clips[i][1])
			i++
		}
		//更新end并计数
		count++
		end = maxEnd
		if end >= time {
			return count
		}
	}
	return -1
}

/*
https://mp.weixin.qq.com/s/hMrwcLn01BpFzBlsvGE2oQ

给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。

贪心思路：
尽可能到达最远位置（贪心）。
如果能到达某个位置，那一定能到达它前面的所有位置。

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

// 一维数组没搞出来
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

/*
	0             7
	2 3 1 1 4 1 1 1

	当前位置i=0, 看能否跳到位置len-1 = 7

	nums[0] = 2
	0 + 1
	0 + 2
	说明当前位置能覆盖的最大位置maxPos是2

	nums[1] = 3(1能被maxPos覆盖)
	1 + 1
	1 + 2
	1 + 3
	说明当前位置能覆盖的最大位置maxPos是4

	nums[2] = 1(2能被maxPos覆盖)
	2 + 1
	说明当前位置能覆盖的最大位置maxPos还是4

	nums[3] = 1(3能被maxPos覆盖)
	3 + 1
	说明当前位置能覆盖的最大位置maxPos还是4

	nums[4] = 4(4能被maxPos覆盖)
	4 + 1
	...
	4 + 4
	说明当前位置能覆盖的最大位置maxPos是8, 已经超过7,ok

   如果延长数组为下面
	0             7   9
	2 3 1 1 4 1 1 1 0 1

	nums[5], nums[6], nums[7] = 1(5,6,7能被maxPos覆盖)
	因为都比8小, 还是取8
	说明当前位置能覆盖的最大位置maxPos是8

	nums[8] = 0(8能被maxPos覆盖)
	8 + 0 = 8
	说明当前位置能覆盖的最大位置maxPos还是8

    nums[9] = 1(9不能被maxPos覆盖, 永远到不了9, 不用计算maxPos了, 结束)


*/
func canJumpGreedy(nums []int) bool {
	maxPos := 0
	for i := 0; i < len(nums); i++ {
		if i > maxPos {
			//当前最远距离maxPos无法超过当前位置i, 说明遇到0卡住了(该0前面所有位置都跳到此0), 无法跳过, 不用计算了
			return false
		}
		//走到这里表示maxPos能覆盖到的地方
		maxPos = Max(maxPos, i+nums[i])
		//这里可以优化下
		//if maxPos >= len(nums) -1 {
		//	return true
		//}
	}
	return maxPos >= len(nums)-1
}

// 45. 跳跃游戏 II
func TestJ(t *testing.T) {
	jumpV2([]int{2, 3, 1, 1, 1, 4, 1, 1, 1})
}

/*
	0 1 2 3 4 5 6 7 8
	2 3 1 1 1 4 1 1 1
      1   3   5 6

这里的思路和上一题一样
    nums[0] = 2, maxPos = 0 + 2 = 2
    nums[1] = 3, maxPos = 1 + 3 = 4
    nums[2] = 1, maxPos = max(2+1, 4) = 4
    nums[3] = 1, maxPos = max(3+1, 4) = 4
    nums[4] = 1, maxPos = max(4+1, 4) = 5
    nums[5] = 4, maxPos = max(5+4, 5) = 9, 已经能跳到了ok, 但是这里不能return, 因为可能后面还有更大的nums[i]产生更小的步数
    nums[6] = 1, maxPos = max(6+1, 9) = 9
    nums[7] = 1, maxPos = max(7+1, 9) = 9


现在的问题就是如何计数, 在什么时机需要计数,

    nums[0] = 2, end = 0, maxPos = 0 + 2 = 2, 索引0能被end覆盖, 不用更新步数
    nums[1] = 3, end = 0, maxPos = 1 + 3 = 4, end = 2, 索引1超过了之前的最大索引0, 设置最大索引为上一个maxPos=2, 新增步数
    nums[2] = 1, end = 2, maxPos = max(2+1, 4) = 4
    nums[3] = 1, end = 2, maxPos = max(3+1, 4) = 4, end = 4, 索引3超过了之前的最大索引2, 设置最大索引为4,  新增步数
    nums[4] = 1, end = 4, maxPos = max(4+1, 4) = 5
    nums[5] = 4, end = 4, maxPos = max(5+4, 5) = 9, end = 5, 索引5超过了之前的最大索引4, 设置最大索引为9,  新增步数         已经能跳到了ok, 但是这里不能return, 因为可能后面还有更大的nums[i]产生更小的步数
    nums[6] = 1, end = 5, maxPos = max(6+1, 9) = 9, end = 9, 索引5超过了之前的最大索引4, 设置最大索引为9,  新增步数
    nums[7] = 1, end = 9, maxPos = max(7+1, 9) = 9


*/

// 题目已保证可以达到
func jumpV2(nums []int) int {
	end := 0    // 当前跳跃范围的边界, 当i超过maxPos时才更新
	maxPos := 0 // 当前能够到达的最远位置
	steps := 0  // 跳跃次数
	for i := 0; i < len(nums); i++ {
		if i > end {
			// 如果当前位置超过了当前跳跃范围的边界，需要增加一次跳跃
			steps++
			end = maxPos
			//fmt.Println(i)
		}
		// 更新当前能够到达的最远位置
		maxPos = Max(maxPos, nums[i]+i)
		fmt.Printf("i := %d, nums[i] = %d, end := %d, maxPos := %d, steps:= %d \n", i, nums[i], end, maxPos, steps)
	}
	return steps
}
