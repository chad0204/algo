package datastruct

import (
	"fmt"
	"testing"
)

/**

1. 舍得用变量，千万别想着节省变量，否则容易被逻辑绕晕
2. head 有可能需要改动时，先增加一个假head(dummy), 返回的时候直接取 dummy.next，这样就不需要为修改 head 增加一大堆逻辑了。


*/

func TestHasCycle(t *testing.T) {

	head := ListNode{1,
		&ListNode{1,
			&ListNode{1,
				&ListNode{1, nil}}}}
	hashCycle(&head)

}

func hashCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func TestDetectCycle(t *testing.T) {

	head := ListNode{0, nil}
	head.Next = &ListNode{1, nil}
	head.Next.Next = &ListNode{2, nil}
	head.Next.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next.Next = &ListNode{4, nil}
	head.Next.Next.Next.Next.Next = &ListNode{5, nil}
	head.Next.Next.Next.Next.Next.Next = &ListNode{6, nil}
	head.Next.Next.Next.Next.Next.Next.Next = &ListNode{7, nil}
	head.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{8, nil}
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{9, nil}
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = head.Next.Next.Next.Next.Next.Next.Next

	fmt.Println(detectCycle(&head))
}

func detectCycle(head *ListNode) *ListNode {
	/*
		求证： a = (n-1)L + c

			设 head到环入口为a, 环入口到相遇点为b, 相遇点到环入口为c
			第一次相遇时, 快指针走的路程为 a + b + n(b + c), 相遇时, 快指针一定已经走了n圈了。(n >= 1)
			任何时候, 快指针都是慢指针的两倍, 第一次相遇时也满足, a + b + n(b + c) = 2(a + b)。（这里有个隐藏条件,a + b能表示慢指针的路程, 说明慢指针进入环后, 只走了b, 没有超过一圈）

			a = n(b+c) + b = (n-1)(b+c) + c, 说明head到环入口的距离a，是相遇到环入口的距离c, 再加上n-1圈数。所以慢指针从head和从相遇点出发, 一定会在入口相遇

			为何慢指针第一圈走不完一定会和快指针相遇？由于快指针比慢指针快一步, 如果都在环中, 那么慢指针静止, 快指针相当于1步的速度前进, 一圈内一定相遇。得出a + b是慢指针在第一次相遇是慢指针的路程

			备注: 此题用map比较普世
	*/

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}

	// 不能直接判断slow == fast , 可能都是nil或者节点只有一个
	if fast == nil || fast.Next == nil {
		return nil
	}

	fast = head
	for fast != head {
		fast = fast.Next
		slow = slow.Next
		//应该先判断再走, 可能相遇点就是起点, 都不会进入循环
		//if fast == slow {
		//	break
		//}
	}
	return slow //fast和slow都可以
}

func TestMergeTwoLists(t *testing.T) {

	l1 := &ListNode{1,
		&ListNode{3, nil}}

	l2 := &ListNode{2,
		&ListNode{4,
			&ListNode{6,
				&ListNode{8, nil}}}}

	fmt.Println(mergeTwoLists(l1, l2))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{-999, nil} //虚拟头节点
	dummy := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			head.Next = l1
			//只有节点符合链表才往前走
			l1 = l1.Next
		} else {
			head.Next = l2
			l2 = l2.Next
		}
		//l1 = l1.Next
		//l2 = l2.Next
		head = head.Next
	}

	if l1 != nil {
		head.Next = l1
	}

	if l2 != nil {
		head.Next = l2
	}
	return dummy.Next
}

// 递归解法
func mergeTwoListsV2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		l1.Next = mergeTwoListsV2(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoListsV2(l1, l2.Next)
		return l2
	}
}

func TestPartition(t *testing.T) {

	head := &ListNode{1,
		&ListNode{5,
			&ListNode{2,
				&ListNode{3,
					&ListNode{4,
						&ListNode{1, nil}}}}}}
	partition(head, 3)

}

// 86. 分隔链表 小于的一个链表, 大于等于的一个链表, 这两个链表原节点顺序不变。 组成一个新链表
func partition(head *ListNode, x int) *ListNode {
	l := &ListNode{-1, nil}
	r := &ListNode{-1, nil}
	dummyL := l
	dummyR := r

	for head != nil {
		if head.Val < x {
			l.Next = head
			l = l.Next
		} else {
			r.Next = head
			r = r.Next
		}
		head = head.Next
	}
	r.Next = nil //不然当x右边有比x小的元素结果有个环

	l.Next = dummyR.Next
	return dummyL.Next
}

func TestMergeKLists(t *testing.T) {

	l1 := &ListNode{1,
		&ListNode{4,
			&ListNode{5, nil}}}

	l2 := &ListNode{1,
		&ListNode{3,
			&ListNode{4, nil}}}

	l3 := &ListNode{2,
		&ListNode{6, nil}}

	//s := make([]*ListNode, 0)
	//s = append(s, l1)
	//s = append(s, l2)
	//s = append(s, l3)

	s := []*ListNode{l1, l2, l3}

	mergeKLists(s)

}

// 合并k个有序链表
func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{-1, nil}
	for i := 0; i < len(lists); i++ {
		head.Next = mergeTwoLists(head.Next, lists[i])
	}
	return head.Next
}

// 倒数第k个节点(从1开始计数)
func getKthFromEnd(head *ListNode, k int) *ListNode {
	f := head
	for k > 0 && f != nil {
		f = f.Next
		k--
	}
	s := head
	if f == nil {
		return s
	}
	for f.Next != nil {
		f = f.Next
		s = s.Next
	}
	return s.Next
}

// 找单链表的中点
func middleNode(head *ListNode) *ListNode {
	s := head
	f := head
	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
	}
	return s
}

func TestGIN(t *testing.T) {
	a := &ListNode{1,
		&ListNode{2, nil}}

	b := &ListNode{11,
		&ListNode{22,
			&ListNode{33,
				&ListNode{44, nil}}}}
	getIntersectionNode(a, b)
}

// 判断两个链表是否相交
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB

	for {
		if a == b {
			//不用担心 a走完headA后, 会继续走headB。 因为a+b = b+a, 如果不相交, 会同时nil
			//要么是相交点、要么是nil
			return a
		}
		if a == nil {
			//当a走完headA, 接着走headB
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			//当b走完headB, 接着走headA
			b = headA
		} else {
			b = b.Next
		}
	}
}

func TestDeleteDuplicates(t *testing.T) {

}

// 83. 删除排序链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	s := head
	f := head

	for f != nil {
		if s.Val == f.Val {
			f = f.Next
		} else {
			s.Next = f
			s = f
		}
	}
	s.Next = nil
	return head
}

// 82. 删除排序链表中的重复元素 II
func deleteDuplicatesV2(head *ListNode) *ListNode {
	//因为head可能被编辑, 所以设置虚拟头节点
	dummy := &ListNode{-1, head}
	s := dummy
	f := head
	for f != nil && f.Next != nil {
		val := f.Val
		if f.Next.Val == val {
			for f != nil && f.Val == val {
				f = f.Next
			}
			s.Next = f
		} else {
			s = f
			f = f.Next
		}
	}

	return dummy.Next
}

/**

[1,1,2,3]

1 1 2 3
s   s s
f f f f f



*/
//19. 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//因为head可能被编辑, 所以设置虚拟头节点
	dummy := &ListNode{-1, head}
	f := dummy
	s := dummy

	for i := 0; i <= n; i++ {
		f = f.Next
	}

	for f != nil {
		s = s.Next
		f = f.Next
	}
	s.Next = s.Next.Next
	return dummy.Next
}

// 206. 反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	//这里不能用newHead, 比如1,2,3,4,5回溯的1时候, head = 1, newHead = 5, 5->1是不对的
	head.Next.Next = head
	head.Next = nil
	return newHead //透传
}

func reverseListIte(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	successor := head

	for successor != nil {
		successor = head.Next
		head.Next = prev

		prev = head
		head = successor
	}
	return prev
}

/**

         2               4
node -> node -> node -> node -> node -> node -> nil

*/

func TestReverseBetween(t *testing.T) {
	head := &ListNode{1,
		&ListNode{2,
			&ListNode{3,
				&ListNode{4,
					&ListNode{5,
						&ListNode{6,
							&ListNode{7, nil}}}}}}}

	between := reverseBetween(head, 2, 4)
	fmt.Println(between)
}

// 92. 反转链表 II
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		node, _ := reverseRight(head, right)
		return node
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

func reverseRight(head *ListNode, right int) (*ListNode, *ListNode) {
	if right == 1 {
		return head, head.Next
	}
	right--
	newHead, successor := reverseRight(head.Next, right)
	head.Next.Next = head
	head.Next = successor
	return newHead, successor
}

/**



node1 -> node2 -> node3 -> node4 -> node5 -> node6 -> node7


  <- node1 <- node2 <- node3   <- node4 <- node5 <- node6  node7 ->

*/

// 25. K 个一组翻转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	//迭代到tail的下一个节点
	successor := head
	idx := 0
	for idx < k {
		if successor == nil {
			//数量不足k, 不翻转
			return head
		}
		successor = successor.Next
		idx++
	}
	//反转head到tail, head做为tail指向successor
	newHead := reverseNode(head, successor)
	//head就是本次反转的tail, tail的next指向下一个的head。递归继续反转successor到k
	head.Next = reverseKGroup(successor, k)
	return newHead
}

func reverseNode(head, successor *ListNode) *ListNode {
	if head.Next == successor {
		return head
	}
	newHead := reverseNode(head.Next, successor)
	head.Next.Next = head
	head.Next = successor
	return newHead
}

func TestName(t *testing.T) {
	head := &ListNode{1,
		&ListNode{2,
			&ListNode{2,
				&ListNode{1, nil}}}}
	isPalindrome(head)
}

// 234. 回文链表
func isPalindrome(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	//处理奇偶
	if fast != nil {
		slow = slow.Next
	}

	left := head
	right := reverseList(slow)

	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}
	return true
}

type Nodee struct {
	Val    int
	Next   *Nodee
	Random *Nodee
}

// 138. 随机链表的复制 (类似的题目 133. 克隆图)
func TestCopy(t *testing.T) {
	head := &Nodee{1, &Nodee{2, &Nodee{3, nil, nil}, nil}, nil}
	list := copyRandomList(head)
	fmt.Println(list)
}

func copyRandomList(head *Nodee) *Nodee {
	exists := make(map[*Nodee]*Nodee)
	return copyRandomListHelper(head, exists)
}

func copyRandomListHelper(head *Nodee, exists map[*Nodee]*Nodee) *Nodee {
	if head == nil {
		return nil
	}

	if _, ok := exists[head]; ok {
		return exists[head]
	}
	newHead := &Nodee{head.Val, nil, nil}
	exists[head] = newHead
	newHead.Next = copyRandomListHelper(head.Next, exists)
	newHead.Random = copyRandomListHelper(head.Random, exists)
	return newHead
}

// 61. 旋转链表
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	//计算长度, 并找到尾节点
	tail := head
	l := 1
	for tail.Next != nil {
		tail = tail.Next
		l++
	}
	k = l - k%l
	if k == l {
		//这里是个小优化
		return head
	}
	//收尾相连
	tail.Next = head
	for k > 1 {
		head = head.Next
		k--
	}
	//断开
	newHead := head.Next
	head.Next = nil
	return newHead
}

func TestSortList(t *testing.T) {
	head := ListNode{-1, nil}
	head.Next = &ListNode{1, nil}
	head.Next.Next = &ListNode{2, nil}
	head.Next.Next.Next = &ListNode{3, nil}
	head.Next.Next.Next.Next = &ListNode{4, nil}
	fmt.Println(sortList(&head))

}

//148. 排序链表
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//slow是真正的mid的前一个
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(mid)
	return mergeTwoLists(left, right)
}
