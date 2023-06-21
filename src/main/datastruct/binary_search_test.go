package datastruct

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 12
	fmt.Println(searchIteration(nums, target))
	fmt.Println(searchRecursion(nums, 0, len(nums)-1, target))
}

// 704. 二分查找
func searchRecursion(nums []int, left, right int, target int) int {
	if left == right {
		if nums[left] == target {
			return left
		} else {
			return -1
		}
	}
	mid := (left + right) / 2
	if nums[mid] > target {
		return searchRecursion(nums, left, mid-1, target)
	} else if nums[mid] < target {
		return searchRecursion(nums, mid+1, right, target)
	} else {
		return mid
	}
}

func searchIteration(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left < right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	if left == right && nums[left] == target {
		return left
	}
	return -1
}

func TestSearchRange(t *testing.T) {
	nums := []int{7, 7, 7, 7, 7, 7}
	fmt.Println(searchRange(nums, 7))
}

/**
0, 1, 2, 3, 4, 5,
5, 7, 7, 7, 8, 10
l     m         r
<
         l  m   r
=
         r/m l
=
      r  l

*/

// 34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	//左边界
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			//搜索区间变成[left, mid-1]
			right = mid - 1
		} else if nums[mid] < target {
			//当left == right时, left向前一步
			//搜索区间变成 [mid+1, right]
			left = mid + 1
		} else {
			//找到target, 收缩右边界, 直到left > right循环退出
			right = mid - 1
		}
	}
	if left < 0 || left >= len(nums) {
		//不合法
		res[0] = -1
	} else if nums[left] == target {
		res[0] = left
	}

	//右边界
	left = 0
	right = len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			//找到相等的 收缩左边界
			left = mid + 1
		}
	}
	if right < 0 || right >= len(nums) {
		//不合法
		res[1] = -1
	} else if nums[right] == target {
		res[1] = right
	}
	return res
}
