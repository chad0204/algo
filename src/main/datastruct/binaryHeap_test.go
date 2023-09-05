package datastruct

import (
	"fmt"
	"testing"
)

/*
完全二叉树： 除最后一层， 其它层都是满的。最后一层靠右。
最大二叉堆： 父节点大于等于左右子节点的完全二叉树
最小二叉堆： 父节点小于等于左右子节点的完全二叉树

	    10
	  9    8
	7  6  5  4
*/

// idx位置上浮到l, 用于新增元素
func minUp(nums []int, idx int, l int) {
	childIdx := idx
	parentIdx := (childIdx - 1) / 2
	value := nums[childIdx]
	for childIdx > l && value < nums[parentIdx] {
		nums[childIdx] = nums[parentIdx]
		childIdx = parentIdx
		parentIdx = (childIdx - 1) / 2
	}
	nums[childIdx] = value
}

func maxUp(nums []int, idx int, l int) {
	childIdx := idx
	parentIdx := (childIdx - 1) / 2
	value := nums[childIdx]
	for childIdx > l && value > nums[parentIdx] {
		nums[childIdx] = nums[parentIdx]
		childIdx = parentIdx
		parentIdx = (childIdx - 1) / 2
	}
	nums[childIdx] = value
}

// idx位置下沉到l, 用于删除元素
func minDown(nums []int, idx int, l int) {
	parentIdx := idx
	childIdx := parentIdx*2 + 1
	value := nums[parentIdx]
	for childIdx < l {
		if childIdx+1 < l && nums[childIdx+1] < nums[childIdx] {
			childIdx = childIdx + 1
		}
		if value < nums[childIdx] {
			break
		}
		nums[parentIdx] = nums[childIdx]
		parentIdx = childIdx
		childIdx = parentIdx*2 + 1
	}
	nums[parentIdx] = value
}

func maxDown(nums []int, idx int, l int) {
	parentIdx := idx
	childIdx := parentIdx*2 + 1
	value := nums[parentIdx]
	for childIdx < l {
		if childIdx+1 < l && nums[childIdx+1] > nums[childIdx] {
			childIdx = childIdx + 1
		}
		if value > nums[childIdx] {
			break
		}
		nums[parentIdx] = nums[childIdx]
		parentIdx = childIdx
		childIdx = parentIdx*2 + 1
	}
	nums[parentIdx] = value
}

func buildMinHeap(nums []int) {
	//最后一个非叶子节点开始, 依次下沉
	lastChildIdx := len(nums)/2 - 1
	for i := lastChildIdx; i >= 0; i-- {
		minDown(nums, i, len(nums))
	}
}

func buildMaxHeap(nums []int) {
	//最后一个非叶子节点开始, 依次上浮
	lastChildIdx := len(nums)/2 - 1
	for i := lastChildIdx; i >= 0; i-- {
		maxDown(nums, i, len(nums))
	}
}

func heapSortASC(nums []int) {
	buildMaxHeap(nums)

	// 堆顶是最大值, 依次把堆顶交换到堆低, 得到从小到大的顺序
	for i := 0; i < len(nums); i++ {
		nums[0], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[0]
		//堆顶下沉
		maxDown(nums, 0, len(nums)-1-i)
	}
}

func heapSortDESC(nums []int) {
	buildMinHeap(nums)

	// 堆顶是最小值, 依次把堆顶交换到堆低, 得到从大到小的顺序
	for i := 0; i < len(nums); i++ {
		nums[0], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[0]
		//堆顶下沉
		minDown(nums, 0, len(nums)-1-i)
	}
}

func TestBinaryHeap(t *testing.T) {
	nums := []int{5, 1, 2, 6, 3, 7, 8, 9, 10, 4}
	//构建
	buildMinHeap(nums)
	fmt.Println(nums)

	fmt.Println("-------插入操作-------")

	//插入尾部
	nums = append(nums, 0)
	fmt.Println("插入尾部: ", nums)
	//上浮调整
	minUp(nums, len(nums)-1, 0)
	fmt.Println("上浮恢复: ", nums)

	fmt.Println("-------删除操作-------")

	//先把头移到尾
	nums[0], nums[len(nums)-1] = nums[len(nums)-1], nums[0]
	fmt.Println("头尾交换: ", nums)
	//删除尾巴
	nums = nums[:len(nums)-1]
	fmt.Println("删除尾部: ", nums)
	//下沉头
	minDown(nums, 0, len(nums)-1)
	fmt.Println("下沉恢复: ", nums)

	fmt.Println("-------排序-------")
	//排序
	heapSortDESC(nums)
	fmt.Println(nums)
}

type PriorityQueue struct {
	size    int
	element []int
}

func buildPriorityQueue(cap int) *PriorityQueue {
	return &PriorityQueue{0, make([]int, cap)}
}

// push 队尾插入元素
func (p *PriorityQueue) push(value int) {
	p.element = append(p.element, value)
	minUp(p.element, len(p.element)-1, 0)
}

// pop 删除队头
func (p *PriorityQueue) pop() int {
	e := p.element[0]
	p.element[0], p.element[len(p.element)-1] = p.element[len(p.element)-1], p.element[0]
	p.element = p.element[:len(p.element)-1]
	minDown(p.element, 0, len(p.element))
	return e
}

func TestPriorityQueue(t *testing.T) {

	pq := buildPriorityQueue(0)

	pq.push(9)
	pq.push(1)
	pq.push(8)
	pq.push(7)
	pq.push(5)
	pq.push(5)
	pq.push(6)
	pq.push(4)
	pq.push(2)
	pq.push(3)
	fmt.Println(pq.element)

	for i := 0; i < 10; i++ {
		fmt.Println(pq.pop())
	}

}

//https://leetcode.cn/problems/kth-largest-element-in-a-stream/?show=1
//https://leetcode.cn/problems/top-k-frequent-words/?show=1
//https://leetcode.cn/problems/sort-characters-by-frequency/?show=1
//https://leetcode.cn/problems/top-k-frequent-elements/description/?show=1
