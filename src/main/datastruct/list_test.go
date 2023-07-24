package datastruct

import (
	"fmt"
	"testing"
)

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
	l := &ListNode{-999, nil} //虚拟头节点
	head := l

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l.Next = l1
			//只有节点符合链表才往前走
			l1 = l1.Next
		} else {
			l.Next = l2
			l2 = l2.Next
		}
		//l1 = l1.Next
		//l2 = l2.Next
		l = l.Next
	}

	if l1 != nil {
		l.Next = l1
	}

	if l2 != nil {
		l.Next = l2
	}

	return head.Next
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
		&ListNode{4,
			&ListNode{3,
				&ListNode{2,
					&ListNode{5,
						&ListNode{2, nil}}}}}}
	partition(head, 3)

}

// 小于的一个链表, 大于等于的一个链表, 这两个链表原节点顺序不变。 组成一个新链表
func partition(head *ListNode, x int) *ListNode {
	l := &ListNode{-1, nil}
	r := &ListNode{-1, nil}

	lhead := l
	rhead := r
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
	l.Next = rhead.Next
	r.Next = nil //不然结果有个环
	return lhead.Next
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
	if len(lists) == 0 {
		return nil
	}

	head := lists[0]
	for i, v := range lists {
		if i == 0 {
			continue
		}
		head = mergeTwoListsV2(head, v)

	}
	return head
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

// 判断两个链表是否相交
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB

	for {
		if a == b {
			//不用担心 a走完headB后, 会继续走headB。 因为a+b = b+a, 如果不相交, 会同时nil
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

// 删除链表中的重复元素
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	s := head
	f := head
	for f != nil {
		if s.Val != f.Val {
			s.Next = f
			s = s.Next
		}
		f = f.Next
	}
	s.Next = nil
	return head
}

/**

[1,1,2,3]

1 1 2 3
s   s s
f f f f f



*/
//19. 删除链表的倒数第 N 个结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{-1, head}
	f := head
	s := dummy

	for i := 0; i < n; i++ {
		f = f.Next
	}

	for f != nil {
		s = s.Next
		f = f.Next
	}
	s.Next = s.Next.Next
	return dummy.Next
}
