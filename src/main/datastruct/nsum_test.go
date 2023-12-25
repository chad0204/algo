package datastruct

import "sort"

/*
https://mp.weixin.qq.com/s/fSyJVvggxHq28a0SdmZm6Q


167. 两数之和 II - 输入有序数组
15. 三数之和
18. 四数之和

*/

// 167. 两数之和 II - 输入有序数组
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		val := numbers[left] + numbers[right]
		if val == target {
			return []int{left + 1, right + 1}
		} else if val < target {
			//nums[left] 加上此时的nums[right]都比target小, 那么left不可能跟right--之后的值组成target了
			left++
		} else {
			right--
		}
	}
	return []int{-1, -1}
}

/**
-1,0,1,2,-1,-4

[-1, 0, 1], [0, 1, -1], 只能取一个

-4 -1 -1 0 1 2

-4  [-1 -1 0 1 2]

-1  [ -1 0 1 2]    [0, 1], [-1, 2]

-1  [0 1 2]        [0, 1]

0


*/
func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		target := -nums[i]
		//这里i+1是避免重复
		twoSums := twoSumNoDup(nums, i+1, target)
		for _, twoSum := range twoSums {
			res = append(res, append(twoSum, nums[i]))
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

func twoSumNoDup(nums []int, start int, target int) [][]int {
	res := make([][]int, 0)
	left := start
	right := len(nums) - 1

	for left < right {
		l := nums[left]
		r := nums[right]
		if l+r == target {
			res = append(res, []int{nums[left], nums[right]})
			//如果相邻相等, 跳过
			for left < right && nums[left] == l {
				left++
			}
			//如果相邻相等, 跳过
			for left < right && nums[right] == r {
				right--
			}
		} else if l+r < target {
			for left < right && nums[left] == l {
				left++
			}
		} else {
			for left < right && nums[right] == r {
				right--
			}
		}
	}
	return res
}
