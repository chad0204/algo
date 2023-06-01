package main

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {

	nums := []int{5, 3, 9, 2, 4, 1, 6, 8, 10}
	quickSort(nums, 0, len(nums)-1)

	fmt.Println(nums)

}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
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

	quickSort(nums, start, l-1)
	quickSort(nums, l+1, end)

}
