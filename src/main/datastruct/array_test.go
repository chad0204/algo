package datastruct

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	a := []int{2, 2, 2, 2, 3, 4, 5, 6, 7}

	removeDuplicates(a)
}

// 删除有序数组中的重复项 (有序是关键) 覆盖！
func removeDuplicates(nums []int) int {
	//2, 2, 2, 2, 3, 4, 5, 6, 7
	//2, 3, 4, 5, 6, 7, 5, 6, 7
	s := 0
	f := 0
	for f < len(nums) {
		if nums[s] != nums[f] {
			s++
			nums[s] = nums[f]
			f++
		} else {
			f++
		}
	}
	return s + 1
}

// 27. 移除元素 此题解法和上题类似
func removeElement(nums []int, val int) int {
	// 1 2 3 2 4 5  删除2
	// 1 3 4 5
	s := 0
	f := 0
	for f < len(nums) {
		if nums[f] != val {
			nums[s] = nums[f]
			f++
			s++
		} else {
			f++
		}
	}
	return s
}

// 两数之和II
func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}
	return []int{-1, -1}
}

// 反转字符串
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		tmp := s[left]
		s[left] = s[right]
		s[right] = tmp
		left++
		right--
	}
}

// 最长回文子串 这题可以动态规划(dp)
func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ { //计算所有节点的中心
		s1 := palindrome(s, i, i)   //奇数中心
		s2 := palindrome(s, i, i+1) //偶数中心
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

func palindrome(s string, l int, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r] //左闭右开
}

// 304. 二维区域和检索 - 矩阵不可变 前缀和
type NumMatrix struct {
	PreNumMatrix [][]int
}

func NumMatrixConstructor(matrix [][]int) NumMatrix {
	row := len(matrix)
	col := len(matrix[0])
	PreNumMatrix := make([][]int, row+1)
	for i := 0; i < len(PreNumMatrix); i++ {
		PreNumMatrix[i] = make([]int, col+1)
	}

	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			fmt.Println(PreNumMatrix[i-1][j], PreNumMatrix[i][j-1], matrix[i-1][j-1], PreNumMatrix[i-1][j-1])
			//x = 上 + 左 + 值 - 对角(多加啦一份重叠的要减掉)
			PreNumMatrix[i][j] = PreNumMatrix[i-1][j] + PreNumMatrix[i][j-1] + matrix[i-1][j-1] - PreNumMatrix[i-1][j-1]
		}
	}
	return NumMatrix{PreNumMatrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	//同理： 多减一份, 要加回来
	return this.PreNumMatrix[row2+1][col2+1] - this.PreNumMatrix[row1][col2+1] - this.PreNumMatrix[row2+1][col1] + this.PreNumMatrix[row1][col1]
}

func TestNumMatrix(t *testing.T) {
	//matrix := make([][]int, 5)
	//matrix[0] = []int{3, 0, 1, 4, 2}
	//matrix[1] = []int{5, 6, 3, 2, 1}
	//matrix[2] = []int{1, 2, 0, 1, 5}
	//matrix[3] = []int{4, 1, 0, 1, 7}
	//matrix[4] = []int{1, 0, 3, 0, 5}

	matrix := make([][]int, 3)
	matrix[0] = []int{1, 2, 3}
	matrix[1] = []int{4, 5, 6}
	matrix[2] = []int{7, 8, 9}

	numMatrix := NumMatrixConstructor(matrix)

	numMatrix.SumRegion(1, 1, 2, 2)
}
