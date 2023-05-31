package datastruct

import "testing"

func TestRemoveDuplicates(t *testing.T) {

	a := []int{1, 2, 2, 2, 3, 4, 5, 6, 7}

	removeDuplicates(a)
}

//删除有序数组中的重复项 (有序是关键)
func removeDuplicates(nums []int) int {
	//0 1 1 1 2 3 4 5 6 7
	//0 1 2 3 4 5 6 7 6 7
	s := 0
	f := 0
	for f < len(nums) {
		if nums[s] != nums[f] {
			s++
			nums[s] = nums[f]
		}
		f++
	}
	return s + 1
}

/**
 0 1 1 1 2 3 4 5 6 7

 0 1 2 3 4 5 6 7 6 7
             s     f
s = 0
f = 1

s = 1
nums[1] = nums[1]
f = 2

s = 1
f = 3

s = 1
f = 4

s = 2
nums[2] = nums[4]
f = 5

s = 3
nums[3] = nums[5]
f = 6




*/
