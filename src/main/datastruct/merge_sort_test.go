package datastruct

/*
归并排序可以看成二叉树的后序遍历

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	left := mergeSort(nums[:len(nums)/2])
	right := mergeSort(nums[len(nums)/2:])
	//合并两个有序数组
	mergeOrderly(left, right)
}

https://www.cnblogs.com/labuladong/p/15943579.html

*/

// 315. 计算右侧小于当前元素的个数
type Pair struct {
	val int
	idx int
}

func countSmaller(nums []int) []int {
	counts := make([]int, len(nums))

	arr := make([]*Pair, len(nums))
	for i := 0; i < len(nums); i++ {
		arr[i] = &Pair{nums[i], i}
	}

	sort(arr, &counts)

	return counts
}

func sort(nums []*Pair, counts *[]int) []*Pair {
	if len(nums) <= 1 {
		return nums
	}
	arr1 := sort(nums[:len(nums)/2], counts)
	arr2 := sort(nums[len(nums)/2:], counts)

	//后序
	return mergeC(arr1, arr2, counts)
}

// 合并两个数组
func mergeC(n1 []*Pair, n2 []*Pair, counts *[]int) []*Pair {
	l := 0
	r := 0
	var tmp []*Pair
	for l < len(n1) && r < len(n2) {
		if n1[l].val <= n2[r].val {
			tmp = append(tmp, n1[l])
			(*counts)[n1[l].idx] += r
			l++
		} else {
			tmp = append(tmp, n2[r])
			r++
		}
	}
	for l < len(n1) {
		tmp = append(tmp, n1[l])
		(*counts)[n1[l].idx] += r
		l++
	}
	for r < len(n2) {
		tmp = append(tmp, n2[r])
		r++
	}
	return tmp
}
