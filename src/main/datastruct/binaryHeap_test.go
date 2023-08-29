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
type BinaryHeap struct {
	heapNums []int
}

type PriorityQueue struct {
	size    int
	element []int
}

// push 队尾插入元素
func (p *PriorityQueue) push(value int) {

}

// pop 删除队头
func (p PriorityQueue) pop() int {
	return 0
}

// idx位置上浮到l, 用于新增元素
func (b *BinaryHeap) swim(idx int, l int) {
	childIdx := idx
	parentIdx := (childIdx - 1) / 2
	value := b.heapNums[childIdx]
	for childIdx > l && value < b.heapNums[parentIdx] {
		b.heapNums[childIdx] = b.heapNums[parentIdx]
		childIdx = parentIdx
		parentIdx = (childIdx - 1) / 2
	}
	b.heapNums[childIdx] = value
}

// idx位置下沉到l, 用于删除元素
func (b *BinaryHeap) minSink(idx int, l int) {
	parentIdx := idx
	childIdx := parentIdx*2 + 1
	value := b.heapNums[parentIdx]
	for childIdx < l {
		if childIdx+1 < l && b.heapNums[childIdx+1] < b.heapNums[childIdx] {
			childIdx = childIdx + 1
		}
		if value < b.heapNums[childIdx] {
			break
		}
		b.heapNums[parentIdx] = b.heapNums[childIdx]
		parentIdx = childIdx
		childIdx = parentIdx*2 + 1
	}
	b.heapNums[parentIdx] = value
}

func (b *BinaryHeap) maxSink(idx int, l int) {
	parentIdx := idx
	childIdx := parentIdx*2 + 1
	value := b.heapNums[parentIdx]
	for childIdx < l {
		if childIdx+1 < l && b.heapNums[childIdx+1] > b.heapNums[childIdx] {
			childIdx = childIdx + 1
		}
		if value > b.heapNums[childIdx] {
			break
		}
		b.heapNums[parentIdx] = b.heapNums[childIdx]
		parentIdx = childIdx
		childIdx = parentIdx*2 + 1
	}
	b.heapNums[parentIdx] = value
}

func buildMinHeap(nums []int) *BinaryHeap {
	b := &BinaryHeap{nums}
	//最后一个非叶子节点开始, 依次下沉
	lastChildIdx := len(nums)/2 - 1
	for i := lastChildIdx; i >= 0; i-- {
		b.minSink(i, len(nums))
	}
	return b
}

func buildMaxHeap(nums []int) *BinaryHeap {
	b := &BinaryHeap{nums}
	//最后一个非叶子节点开始, 依次上浮
	lastChildIdx := len(nums)/2 - 1
	for i := lastChildIdx; i >= 0; i-- {
		b.maxSink(i, len(nums))
	}
	return b
}

func heapSortASC(nums []int) {
	heap := buildMaxHeap(nums)

	// 堆顶是最大值, 依次把堆顶交换到堆低, 得到从小到大的顺序
	for i := 0; i < len(nums); i++ {
		heap.heapNums[0], heap.heapNums[len(nums)-1-i] = heap.heapNums[len(nums)-1-i], heap.heapNums[0]
		//堆顶下沉
		heap.maxSink(0, len(nums)-1-i)
	}
}

func heapSortDESC(nums []int) {
	heap := buildMaxHeap(nums)

	// 堆顶是最大值, 依次把堆顶交换到堆低, 得到从小到大的顺序
	for i := 0; i < len(nums); i++ {
		heap.heapNums[0], heap.heapNums[len(nums)-1-i] = heap.heapNums[len(nums)-1-i], heap.heapNums[0]
		//堆顶下沉
		heap.maxSink(0, len(nums)-1-i)
	}
}

func TestBinaryHeap(t *testing.T) {
	nums := []int{5, 1, 2, 6, 3, 7, 8, 9, 10}
	//构建
	heap := buildMinHeap(nums)
	fmt.Println(heap.heapNums)

	fmt.Println("-------插入操作-------")

	//插入尾部
	heap.heapNums = append(heap.heapNums, 0)
	fmt.Println(heap.heapNums)
	//上浮调整
	heap.swim(len(heap.heapNums)-1, 0)
	fmt.Println(heap.heapNums)

	fmt.Println("-------删除操作-------")

	//先把头移到尾
	heap.heapNums[0], heap.heapNums[len(heap.heapNums)-1] = heap.heapNums[len(heap.heapNums)-1], heap.heapNums[0]
	fmt.Println(heap.heapNums)
	//删除尾巴
	heap.heapNums = heap.heapNums[:len(heap.heapNums)-1]
	fmt.Println(heap.heapNums)
	//下沉头
	heap.minSink(0, len(heap.heapNums)-1)
	fmt.Println(heap.heapNums)

	fmt.Println("-------排序-------")
	//排序
	heapSortASC(heap.heapNums)
	fmt.Println(heap.heapNums)
}
