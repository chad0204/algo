package datastruct

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	a := []int{2, 2, 2, 2, 3, 4, 5, 6, 7}

	removeDuplicates(a)
}

// 删除有序数组中的重复项 (有序是关键, 有序表示重复的一定相连) 覆盖！
func removeDuplicates(nums []int) int {
	//2, 2, 2, 2, 3, 4, 5, 6, 7
	//2, 3, 4, 5, 6, 7, 5, 6, 7
	s := 0
	f := 0
	for f < len(nums) {
		if nums[s] != nums[f] {
			s++
			nums[s] = nums[f]
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

			//x = 上 + 左 + 值 - 对角(多加啦一份重叠的要减掉)
			PreNumMatrix[i][j] = PreNumMatrix[i-1][j] + PreNumMatrix[i][j-1] + matrix[i-1][j-1] - PreNumMatrix[i-1][j-1]
			fmt.Print(PreNumMatrix[i][j])
		}
		fmt.Println()
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

	fmt.Println(numMatrix.SumRegion(0, 0, 2, 2))

}

// 1094. 拼车 差分数组
func TestCarPooling(t *testing.T) {
	//[[2,1,5],[3,5,7]] 3
	trips := make([][]int, 2)
	trips[0] = []int{2, 1, 5}
	trips[1] = []int{3, 5, 7}
	carPooling(trips, 3)
}

func carPooling(trips [][]int, capacity int) bool {
	nums := make([]int, 1001)

	diff := NewDiffNums(nums)
	for _, v := range trips {
		Increment(diff, v[0], v[1], v[2]-1) // -1 是因为下车和上车同一个位置是不算总数的
	}
	res := GetRes(diff)

	for _, v := range res {
		if v > capacity {
			return false
		}
	}
	return true
}

func Increment(diff []int, val int, i, j int) {
	diff[i] += val
	if j+1 < len(diff) {
		diff[j+1] -= val
	}
}

func NewDiffNums(nums []int) []int {
	diff := make([]int, len(nums))
	diff[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diff[i] = nums[i] - nums[i-1]
	}
	return diff
}

func GetRes(diff []int) []int {
	nums := make([]int, len(diff))
	nums[0] = diff[0]
	for i := 1; i < len(diff); i++ {
		nums[i] = diff[i] + nums[i-1]
	}
	return nums
}

// 1109. 航班预订统计 差分数组
func corpFlightBookings(bookings [][]int, n int) []int {
	nums := make([]int, n)
	diff := NewDiff(nums)

	for i, _ := range bookings {
		diff.Incr(bookings[i][0]-1, bookings[i][1]-1, bookings[i][2])
	}
	return diff.GetRes()
}

type Diff struct {
	diffNums []int
}

func NewDiff(nums []int) *Diff {
	diffNums := make([]int, len(nums))
	diffNums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		diffNums[i] = nums[i] - diffNums[i-1]
	}
	return &Diff{diffNums}
}

func (d *Diff) GetRes() []int {
	nums := make([]int, len(d.diffNums))
	nums[0] = d.diffNums[0]
	for i := 1; i < len(d.diffNums); i++ {
		nums[i] = d.diffNums[i] + nums[i-1]
	}
	return nums
}

func (d *Diff) Incr(i, j, val int) {
	d.diffNums[i] += val
	if j+1 < len(d.diffNums) {
		d.diffNums[j+1] -= val
	}
}

// 48. 旋转图像 先对角线翻转, 再逐行翻转
func rotate(matrix [][]int) {
	n := len(matrix)
	// 沿对角线镜像反转, 只能遍历半个矩阵, 不然就转回去了
	for row := 0; row < n; row++ {
		for col := row; col < n; col++ {
			matrix[row][col], matrix[col][row] = matrix[col][row], matrix[row][col]
		}
	}
	//逐行反转
	for i, _ := range matrix {
		reverse(matrix[i])
	}
}

func reverse(nums []int) {
	l := 0
	r := len(nums) - 1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

// 54. 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	n := len(matrix[0])
	res := make([]int, 0, m*n) // 就是长度为0, cap为m*n的数组

	left := 0      //左边界++
	right := n - 1 //右边界--
	up := 0        //上边界++
	low := m - 1   //下边界--

	for up <= low && left <= right {
		if up <= low { //表示up到low之间还有行
			for i := left; i <= right; i++ {
				res = append(res, matrix[up][i])
			}
			//只有遍历过这行, 才能++
			up++
		}

		if left <= right { //表示left到right之间还有列
			for i := up; i <= low; i++ {
				res = append(res, matrix[i][right])
			}
			//只有遍历过这列, 才能++
			right--
		}

		if up <= low {
			for i := right; i >= left; i-- {
				res = append(res, matrix[low][i])
			}
			low--
		}

		if left <= right {
			for i := low; i >= up; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}

	}
	return res
}

/**
36. 有效的数独

board[i][j]
所在的行是下标i, 所在的列是下j, 所在的box是x
条件：
1. board[i][j]在i行只出现一次
2. board[i][j]在j列只出现一次
3. board[i][j]在box[x]只出现一次


如何根据i,j定位box（9x9的矩阵, box有9个，分别为0到8）：

0 1 2
3 4 5
6 7 8

1. j找列
j/3表示所在的box为第几列(0,1,2), 不管i
0/3 = 0
1/3 = 0
2/3 = 0
..
4/3 = 1
..
8/3 = 2


2. i找行, 前面判断列, 算上i的话, 如果i在012, 就是box就是j/3, 不用加; 如果i在345, box需要加1行就是3; 如果i在678, box需要加2行就是6; 也就是 j/3 + (i/3) * 3
i/3表示所在的box为第几行(0,1,2)
0/3 = 0
1/3 = 0
2/3 = 0
..
4/3 = 1
..
8/3 = 2
*/

type Pos struct {
	pos int
	val byte
}

func isValidSudoku(board [][]byte) bool {
	// board[i][j] 在第i行是否存在
	rowMap := make(map[Pos]bool)
	colMap := make(map[Pos]bool)
	boxMap := make(map[Pos]bool)

	//也可以这么定义
	//row := make([][10]bool, 9)   // 存储每一行的每个数是否出现过, 9行10个数
	//col := make([][10]bool, 9)   // 存储每一列的每个数是否出现过
	//box := make([][10]bool, 9)   // 存储每一个box的每个数是否出现过

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			val := board[i][j]
			if val == '.' {
				continue
			}
			if rowMap[Pos{i, val}] {
				return false
			}
			if colMap[Pos{i, val}] {
				return false
			}
			if boxMap[Pos{i, val}] {
				return false
			}

			rowMap[Pos{i, val}] = true
			colMap[Pos{j, val}] = true
			boxMap[Pos{j/3 + (i/3)*3, val}] = true

		}
	}
	return true
}
