package datastruct

import (
	"fmt"
	"math"
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
	//上面left < right 改成 left <= right, 这段可以去掉
	if left == right && nums[left] == target {
		return left
	}
	return -1
}

func TestSearchRange(t *testing.T) {
	nums := []int{7, 7, 7, 7, 7, 7}
	fmt.Println(searchRange(nums, 7))
}

// 34. 在排序数组中查找元素的第一个和最后一个位置
func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			//[mid, right]找右极限
			left = mid + 1
		} else if nums[mid] > target {
			//[left, mid]
			right = mid - 1
		} else {
			//nums[mid] = target, 左边界从mid递减, 右边界从mid递增
			left, right = mid, mid
			for left >= 0 {
				if nums[left] != target {
					break
				}
				left--
			}
			for right < len(nums) {
				if nums[right] != target {
					break
				}
				right++
			}
			return []int{left + 1, right - 1}
		}
	}
	return []int{-1, -1}
}

func TestFPE(t *testing.T) {
	nums := []int{2, 3, 4, 5, 1}
	fmt.Print(findMin(nums))

}

// 162. 寻找峰值, 虽然题目要求只要符合左右相邻元素都比当前小的就满足, 可以转成找“局部”最大值(这里的最大值不是全局最大值), 因为局部最大值一定满足(题目要求相邻不想等)
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	//注意这里不能等于left == right就是结果, 不能继续
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			//下降阶段, 最大值在[left:mid]
			right = mid
		} else {
			//上升阶段, 最大值在(mid:right]
			left = mid + 1
		}
	}
	return left
}

/**
通用的寻找最大值和最小值
func findMin(nums []int) int {
    idx := 0
    for i,v := range nums {
        if v < nums[idx] {
            idx = i
        }
    }
    return nums[idx]
}

func findMax(nums []int) int {
    idx := 0
    for i,v := range nums {
        if v > nums[idx] {
            idx = i
        }
    }
    return nums[idx]
}
*/

// 153. 寻找旋转排序数组中的最小值
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	res := math.MaxInt32
	for left <= right {
		mid := left + (right-left)/2
		if nums[left] <= nums[mid] {
			//左边有序
			res = Min(res, nums[left])
			left = mid + 1
		} else {
			res = Min(res, nums[mid])
			right = mid - 1
		}
	}
	return res
}

// 33. 搜索旋转排序数组
func search(nums []int, target int) int {
	//二分查找最大值
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		/*
				肯定有一半有有序的
				这里有个边界, nums[left] == nums[mid]的情况,
				nums[left] == nums[mid]表示(l+r)/2 = mid = l, 有两种情况：1. r=l; 2.r=l+1

			如果l==r且不是target, 随便哪边都行, 最后会返回-1
			如果l+1==r且不是target, 只能试试left = mid+1
		*/

		if nums[left] <= nums[mid] {
			//左半边有序
			if nums[left] <= target && target < nums[mid] {
				//在左区间
				right = mid - 1
			} else {
				//不在左区间, 朝右边逼近
				left = mid + 1
			}
		} else {
			//右半边有序
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				//不在右区间, 朝左边逼近
				right = mid - 1
			}
		}
	}
	return -1
}
