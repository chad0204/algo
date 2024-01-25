package datastruct

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

/*
*

多叉树遍历

回溯算法, 遍历树枝
DFS, 遍历节点

标准: 排列(不在乎顺序, 无间隔, 121和112都是排列) -> 组合 -> 子集(在乎顺序, 可以间隔, 12和21是一样的只能取一个)

变型: n皇后
变型: 有重复数字的排列, 不可复选
变型: 有重复数字的子集, 不可复选

i := start 同级不能回头, 从上一层递归的索引开始遍历
used 只能选一次
!used[i-1] && nums[i-1] == nums[i] 重复值
*/
func TestPermute(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3}))
}

// 46. 全排列
func permute(nums []int) [][]int {
	permutes = [][]int{}            //每次执行都清空结果, 或者backtrack传递指针参数也行
	used := make([]bool, len(nums)) // 记录一次深度遍历的已使用索引
	var path []int                  //不要初始值                //不能使用make, make不仅分配内存还会设置初始值
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

func TestSolveNQueens(t *testing.T) {
	fmt.Println(solveNQueens(4))
}

// 51. N 皇后 用递归遍历行 迭代遍历列
func solveNQueens(n int) [][]string {
	solves = [][]string{}
	// 每个元素表示每一行的放置情况, 凑齐所有行数就是一个解决方案
	board := make([]string, n)
	//{"....", "....", "....", "...."}
	for i := range board {
		board[i] = strings.Repeat(".", n)
	}
	row := 0 // 从第0行开始
	backtrackBoard(board, row)
	return solves
}

// solves := [][]string{[]string{"Q..."}, []string{".Q.."}}
var solves [][]string

// 由于一行设置完就不能在同一行继续设置 所以从上往下一行一行设置
func backtrackBoard(board []string, row int) {
	if row == len(board) {
		//已设置到最后一行 说明是一种解决方案
		s := make([]string, len(board))
		copy(s, board)
		solves = append(solves, s)
		return
	}

	//遍历n行n列
	for col := 0; col < len(board[row]); col++ {
		//判断此位置能否设置Q
		if !isValid(board, row, col) {
			continue
		}
		rowLine := []byte(board[row])
		rowLine[col] = 'Q'
		board[row] = string(rowLine)
		backtrackBoard(board, row+1)
		rowLine[col] = '.'
		board[row] = string(rowLine)
	}

}

/*
*
判断（row, col) 位置在board上是否可行
*/
func isValid(board []string, row int, col int) bool {
	//列
	for i := row - 1; i >= 0; i-- {
		if []byte(board[i])[col] == 'Q' {
			return false
		}
	}

	//右上
	for i, j := row-1, col+1; i >= 0 && j < len(board[i]); i, j = i-1, j+1 {
		if []byte(board[i])[j] == 'Q' {
			return false
		}
	}

	//左上
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if []byte(board[i])[j] == 'Q' {
			return false
		}
	}
	return true
}

func TestPermuteUnique(t *testing.T) {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}

// 47. 全排列 II (nums包含重复数字)
func permuteUnique(nums []int) [][]int {
	var r [][]int
	used := make([]bool, len(nums))
	var path []int
	sort.Ints(nums)
	backTrackUnique(nums, path, used, &r)
	return r
}

func backTrackUnique(nums []int, path []int, used []bool, r *[][]int) {
	if len(path) == len(nums) {
		s := make([]int, len(path))
		copy(s, path)
		*r = append(*r, s)
		return
	}

	for i := range nums {
		//深度遍历是否有过这个分支
		if used[i] {
			continue
		}
		//与前一个相邻分支值重复 并且前一个相邻分支没有在深度遍历中使用, 则跳过
		if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
			// 如果used[i-1] == true, 说明相邻节点已经在深度树中, 不应该剪切当前分支(前提是排序)
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		backTrackUnique(nums, path, used, r)
		path = path[:len(path)-1]
		used[i] = false
	}
}

func backTrackUniqueV2(nums []int, path []int, used []bool, r *[][]int) {
	if len(path) == len(nums) {
		s := make([]int, len(path))
		copy(s, path)
		*r = append(*r, s)
		return
	}

	pre := -999
	for i := range nums {
		//深度遍历是否有过这个分支
		if used[i] {
			continue
		}
		if nums[i] == pre {
			//pre只有在遍历上一个分支没有会被剪掉的时候才会记录
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		pre = nums[i]
		backTrackUnique(nums, path, used, r)
		path = path[:len(path)-1]
		used[i] = false
	}
}

func TestSubsets(t *testing.T) {
	fmt.Println(subsets([]int{1, 2, 3}))
}

// 78. 子集
/*
         []
      /    \    \
    [1]     [2] [3]
    /   \     \
 [1,2] [1,3]  [3]
  /
[1,2,3]
*/
func subsets(nums []int) [][]int {
	var path []int
	used := make([]bool, len(nums)) // 想想怎么可以不用标记使用
	var res [][]int
	backTrackSubsets(nums, path, 0, used, &res)
	return res
}

func backTrackSubsets(nums []int, path []int, index int, used []bool, res *[][]int) {
	s := make([]int, len(path))
	copy(s, path)
	*res = append(*res, s)
	for i := index; i < len(nums); i++ {
		if used[i] {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		backTrackSubsets(nums, path, i, used, res)
		path = path[:len(path)-1]
		used[i] = false
	}
}

func TestCombine(t *testing.T) {
	fmt.Println(combine(4, 2))
}

// 77. 组合
func combine(n int, k int) [][]int {
	var nums []int // 想想怎么可以不用构造
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	var path []int
	used := make([]bool, n) //想想怎么可以不用标记使用, backtrackCombine(nums, path, i+1, k, used, res)， 这里取i+1即可
	var res [][]int
	backtrackCombine(nums, path, 0, k, used, &res)
	return res
}

func backtrackCombine(nums []int, path []int, index int, k int, used []bool, res *[][]int) {
	if len(path) == k {
		s := make([]int, k)
		copy(s, path)
		*res = append(*res, s)
		return
	}
	for i := index; i < len(nums); i++ {
		if used[i] {
			continue
		}
		path = append(path, nums[i])
		used[i] = true
		backtrackCombine(nums, path, i, k, used, res)
		path = path[:len(path)-1]
		used[i] = false
	}
}

func TestSubsetsWithDup(t *testing.T) {
	fmt.Println(subsetsWithDup([]int{1, 1, 2}))
}

// 90. 子集 II(有重复数据)
func subsetsWithDup(nums []int) [][]int {
	var path []int
	var res [][]int
	sort.Ints(nums)
	backtrackSubsetsWithDup(nums, path, 0, &res)
	return res
}

func backtrackSubsetsWithDup(nums []int, path []int, index int, res *[][]int) {
	s := make([]int, len(path))
	copy(s, path)
	*res = append(*res, s)
	for i := index; i < len(nums); i++ {
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		backtrackSubsetsWithDup(nums, path, i+1, res)
		path = path[:len(path)-1]
	}
}

// 39. 组合总和
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	backtrackCombinationSum(candidates, path, 0, 0, target, &res)
	return res
}

func backtrackCombinationSum(candidates []int, path []int, sum int, start int, target int, res *[][]int) {
	if sum > target {
		return
	}
	if sum == target {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*res = append(*res, tmp)
	}
	for i := start; i < len(candidates); i++ {
		path = append(path, candidates[i])
		sum = sum + candidates[i]
		backtrackCombinationSum(candidates, path, sum, i, target, res)
		path = path[:len(path)-1]
		sum = sum - candidates[i]
	}
}

func TestCombinationSum3(t *testing.T) {
	combinationSum3(3, 7)
}

// 216. 组合总和 III
func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var nums []int
	var path []int
	used := make([]bool, 9)
	for i := 1; i <= 9; i++ {
		nums = append(nums, i)
	}
	backtrackCombinationSum3(nums, path, 0, used, 0, k, n, &res)
	return res
}

func backtrackCombinationSum3(nums []int, path []int, sum int, used []bool, start int, k int, n int, res *[][]int) {
	if sum > n {
		return
	}
	if len(path) > k {
		return
	}
	if sum == n && len(path) == k {
		s := make([]int, len(path))
		copy(s, path)
		*res = append(*res, s)
		return
	}
	for i := start; i < len(nums); i++ {
		if used[i] {
			continue
		}
		path = append(path, nums[i])
		sum = sum + nums[i]
		used[i] = true
		backtrackCombinationSum3(nums, path, sum, used, i, k, n, res)
		path = path[:len(path)-1]
		sum = sum - nums[i]
		used[i] = false
	}
}

func TestCombinationSum2(t *testing.T) {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

// 40. 组合总和 II
func combinationSum2(candidates []int, target int) [][]int {
	var path []int
	used := make([]bool, len(candidates)) //可以去掉
	var res [][]int
	sort.Ints(candidates)
	backtrackCombinationSum2(candidates, path, used, 0, 0, target, &res)
	return res
}

func backtrackCombinationSum2(candidates []int, path []int, used []bool, start int, sum int, target int, res *[][]int) {
	if sum > target {
		return
	}
	if sum == target {
		s := make([]int, len(path))
		copy(s, path)
		*res = append(*res, s)
	}
	for i := start; i < len(candidates); i++ {
		if used[i] {
			continue
		}
		if i > 0 && (candidates[i] == candidates[i-1]) && !used[i-1] {
			continue
		}
		path = append(path, candidates[i])
		sum = sum + candidates[i]
		used[i] = true
		backtrackCombinationSum2(candidates, path, used, i, sum, target, res)
		path = path[:len(path)-1]
		sum = sum - candidates[i]
		used[i] = false
	}
}
