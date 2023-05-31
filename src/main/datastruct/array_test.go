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

//删除指定元素 此题解法和上题类似
func removeElement(nums []int, val int) int {
	// 1 2 3 2 4 5  删除2
	// 1 3 4 5
	s := 0
	f := 0
	for f < len(nums) {
		if nums[f] != val {
			nums[s] = nums[f]
			s++
		}
		f++
	}
	return s
}
