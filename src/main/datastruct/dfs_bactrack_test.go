package datastruct

import (
	"fmt"
	"strings"
	"testing"
)

/*
*

多叉树遍历

回溯算法, 遍历树枝
DFS, 遍历节点
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

// 51. N 皇后
func solveNQueens(n int) [][]string {
	solves = [][]string{}
	// 每个元素表示每一行的放置情况, 凑齐所有行数就是一个解决方案
	board := make([]string, n)
	//board = []string{"....", "....", ".....", "....."}
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
	for i := row; i > 0; i-- {
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
