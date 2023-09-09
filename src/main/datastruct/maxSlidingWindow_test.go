package datastruct

import (
	"fmt"
	"testing"
)

//239. 滑动窗口最大值

/**

输入：nums = [1,3,9,-3,5,3,6,7], k = 3

i 先走3步
i      3 4 5 6 7
1,8,9,-3,5,3,6,7
i < 3, 最值为9
i = 3, 最值为9
i = 4, 最值为9
i = 5, 堆中最大的还是9, 但是9已经不在窗口[-3, 5, 3]内,所以把9弹出, 最值为8, 还是不在, 弹出8, 最值为5
i = 6, 最值为6,
i = 7, 最值为7
*/

func TestMSW(t *testing.T) {
	fmt.Println(maxSlidingWindowV2([]int{1, 3, 9, -3, 5, 3, 6, 7}, 3))
}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	priorityQueue := &PQ{make([][]int, 0), 0}
	res := make([]int, 0)
	//先走k步, 并取出前k个元素的最大值
	for i := 0; i < k; i++ {
		priorityQueue.offer([]int{nums[i], i})
	}
	res = append(res, priorityQueue.peek()[0])
	for i := k; i < n; i++ {
		priorityQueue.offer([]int{nums[i], i})
		//此时的最大值如果不在滑动窗口中, 弹出去
		for priorityQueue.peek()[1] <= i-k {
			priorityQueue.poll()
		}
		res = append(res, priorityQueue.peek()[0])
	}
	return res
}

type PQ struct {
	nums [][]int
	size int
}

func (p *PQ) poll() []int {
	value := p.nums[0]
	p.nums[0], p.nums[p.size-1] = p.nums[p.size-1], p.nums[0]
	p.nums = p.nums[:p.size-1]
	p.size--
	down(p.nums, 0, p.size)
	return value
}

func (p *PQ) peek() []int {
	return p.nums[0]
}

func (p *PQ) offer(value []int) {
	p.nums = append(p.nums, value)
	up(p.nums, p.size, 0)
	p.size++
}

func up(nums [][]int, bottom int, targetIdx int) {
	childIdx := bottom
	parentIdx := (childIdx - 1) / 2
	value := nums[bottom]
	for childIdx > targetIdx && value[0] > nums[parentIdx][0] {
		nums[childIdx] = nums[parentIdx]
		childIdx = parentIdx
		parentIdx = (childIdx - 1) / 2
	}
	nums[childIdx] = value
}

func down(nums [][]int, top int, targetIdx int) {
	if len(nums) == 0 {
		return
	}
	parentIdx := top
	childIdx := parentIdx*2 + 1
	value := nums[top]
	for childIdx < targetIdx {
		if childIdx+1 < targetIdx && nums[childIdx+1][0] > nums[childIdx][0] {
			childIdx = childIdx + 1
		}
		if value[0] > nums[childIdx][0] {
			break
		}
		nums[parentIdx] = nums[childIdx]
		parentIdx = childIdx
		childIdx = parentIdx*2 + 1
	}
	nums[parentIdx] = value
}

/**
单调队列


输入：nums = [1,3,9,-3,5,3,6,7], k = 3

1,3,9,-3,5,3,6,7



*/
func maxSlidingWindowV2(nums []int, k int) []int {
	n := len(nums)
	q := &Queue{make([]int, 0)}
	res := make([]int, 0)

	for i := 0; i < n; i++ {
		q.offer(nums[i])
		if i < k-1 {
			continue
		}
		res = append(res, q.max())
		q.poll(nums[i-k+1])
	}
	return res
}

type Queue struct {
	nums []int
}

func (p *Queue) offer(value int) {
	for len(p.nums) > 0 && p.nums[len(p.nums)-1] < value {
		p.nums = p.nums[:len(p.nums)-1]
	}
	p.nums = append(p.nums, value)
}

func (p *Queue) poll(value int) {
	if value == p.nums[0] {
		p.nums = p.nums[1:]
	}
}

func (p *Queue) max() int {
	return p.nums[0]
}
