package datastruct

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	nums := []int{5, 3, 9, 2, 4, 1, 6, 8, 10}
	quickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func quickSort(nums []int, start, end int) {
	if start > end {
		return
	}
	//前序
	partition := part(nums, start, end)
	quickSort(nums, start, partition-1)
	quickSort(nums, partition+1, end)
}

func part(nums []int, start, end int) int {
	//第一位, 最后一位, 中间值, 随机值
	p := nums[start]
	l := start
	r := end
	for l < r {
		for nums[r] >= p && l < r {
			r--
		}
		for nums[l] <= p && l < r {
			l++
		}
		//go 直接交换
		nums[l], nums[r] = nums[r], nums[l]
	}
	nums[start], nums[l] = nums[l], nums[start]
	return l
}

func TestMergeSort(t *testing.T) {
	sort := mergeSort([]int{5, 3, 9, 2, 4, 1, 6, 8, 10})
	fmt.Println(sort)
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	arr1 := mergeSort(nums[:len(nums)/2])
	arr2 := mergeSort(nums[len(nums)/2:])

	//todo 优化 如果arr1 < arr2, 直接return

	//后序
	return merge(arr1, arr2)
}

// 合并两个数组
func merge(n1 []int, n2 []int) []int {
	l := 0
	r := 0
	var tmp []int
	for l < len(n1) && r < len(n2) {
		if n1[l] <= n2[r] {
			tmp = append(tmp, n1[l])
			l++
		} else {
			tmp = append(tmp, n2[r])
			r++
		}
	}
	for l < len(n1) {
		tmp = append(tmp, n1[l])
		l++
	}
	for r < len(n2) {
		tmp = append(tmp, n2[r])
		r++
	}
	return tmp
}

func TestBubbleSort(t *testing.T) {
	nums := []int{5, 3, 9, 2, 4, 1, 6, 8, 10}
	bubbleSort(nums)
	fmt.Println(nums)
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func TestHeapSort(t *testing.T) {
	nums := []int{5, 3, 9, 2, 4, 1, 6, 8, 10}
	heapSort(nums)
	fmt.Println(nums)
}

func heapSort(nums []int) {
	heapify(nums)
	n := len(nums) - 1
	for i := 0; i <= n; i++ {
		nums[0], nums[n-i] = nums[n-i], nums[0]
		heapifyDown(nums, 0, n-i)
	}
}

func heapify(nums []int) {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		heapifyDown(nums, i, n)
	}
}

func heapifyDown(nums []int, lastParentIdx int, n int) {
	childIdx := lastParentIdx*2 + 1
	for childIdx < n {
		if childIdx+1 < n && nums[childIdx] < nums[childIdx+1] {
			//存在比左子节点大的右子节点
			childIdx++
		}
		//不用下沉了
		if nums[lastParentIdx] >= nums[childIdx] {
			break
		}
		nums[lastParentIdx], nums[childIdx] = nums[childIdx], nums[lastParentIdx]
		lastParentIdx = childIdx
		childIdx = lastParentIdx*2 + 1
	}
}
