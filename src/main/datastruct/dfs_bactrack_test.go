package datastruct

import "testing"

/*
*

多叉树遍历

回溯算法, 遍历树枝
DFS, 遍历节点
*/
func TestPermute(t *testing.T) {

	permute([]int{1, 2, 3})
}

// 46. 全排列
func permute(nums []int) [][]int {
	permutes = [][]int{} //每次执行都清空结果, 或者backtrack传递指针参数也行
	used := make([]bool, len(nums))
	var path []int //不能使用make, make不仅分配内存还会设置初始值
	backtrack(nums, path, used)
	return permutes
}

var permutes [][]int

func backtrack(nums []int, path []int, used []bool) {
	if len(nums) == len(path) {
		s := make([]int, len(path))
		copy(s, path)
		permutes = append(permutes, s)
		return
	}
	for i := range nums {
		if used[i] {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		backtrack(nums, path, used)
		path = path[:len(path)-1]
		used[i] = false
	}
}
