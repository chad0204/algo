package datastruct

import (
	"fmt"
	"testing"
)

func TestNextGE(t *testing.T) {
	fmt.Print(nextGreaterElements([]int{1, 2, 3, 4, 5, 6, 5, 4, 5, 1, 2, 3}))
}

// 496. 下一个更大元素 I
func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	s := make([]int, 0)
	for i := len(nums) - 1; i >= 0; i-- {
		for len(s) > 0 && s[len(s)-1] <= nums[i] {
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			idx := -1
			for j := 0; j < i; j++ {
				if nums[j] > nums[i] {
					idx = j
					break
				}
			}
			if idx == -1 {
				res[i] = -1
			} else {
				res[i] = nums[idx]
			}
		} else {
			res[i] = s[len(s)-1]
		}
		s = append(s, nums[i])
	}
	return res
}
