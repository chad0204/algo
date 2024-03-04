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
	n := len(nums)
	res := []int{-1, -1}
	//寻找左边界
	l, r := 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if target < nums[mid] {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	if l < n && nums[l] == target {
		//防止nums任意数都比target大, 导致l溢出
		res[0] = l
	}

	//寻找右边界
	l, r = 0, n-1
	for l <= r {
		mid := l + (r-l)/2
		if target < nums[mid] {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			l = mid + 1
		}
	}
	if r >= 0 && nums[r] == target {
		//防止nums任意数都比target小, 导致r为负数
		res[1] = r
	}
	return res
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

// 153. 寻找旋转排序数组中的最小值 更好理解
func findMinV2(nums []int) int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		if l == r {
			break
		}
		mid := l + (r-l)/2
		if nums[mid] > nums[r] {
			//一定在右边
			l = mid + 1
		} else {
			//在mid左边, 或者等于mid
			r = mid
		}
	}
	return nums[l]
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

			nums[left] <= nums[mid] , 算左区间有序
			nums[mid] <= nums[right], 算右区间有序

		*/

		if nums[left] < nums[mid] {
			//左半边有序
			if nums[left] <= target && target < nums[mid] {
				//在左区间
				right = mid - 1
			} else {
				//不在左区间, 朝右边逼近
				left = mid + 1
			}
		} else if nums[mid] < nums[right] {
			//右半边有序
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				//不在右区间, 朝左边逼近
				right = mid - 1
			}
		} else if nums[left] == nums[mid] {
			//相等说明left==right  or left+1 == right, 也能遇到重复值, 能走到这说明nums[left]不是, left = mid + 1
			left = mid + 1
		} else if nums[mid] == nums[right] {
			//相等说明left==right, 能走到这说明nums[right]不是, 这里可以用mid-1, 也可以直接返回-1
			right = mid - 1
		}
	}
	return -1
}
