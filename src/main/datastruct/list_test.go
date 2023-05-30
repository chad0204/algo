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

	head := ListNode{1, nil}
	head.Next = &ListNode{2, &head}

	fmt.Println(detectCycle(&head))
}

func detectCycle(head *ListNode) *ListNode {
	/*
		设 环一圈长度为L
		第一次相遇时, 慢指针走了k步, 快指针走了2k步, 由于相遇时, 快指针已经在环中转了n圈, 所以k = nL
		设相遇点距离环入口距离是M步

		此时慢指针从head走(K-M)步就是环入口。
		慢指针从相遇点走多少步能到环入口呢？ (L-M) + (n-1)L = nL - M = K-M ,（L-M到环入口的第一圈, n-1是后续圈数）所以 慢指针从相遇点也是走K-M步到环起点

		第一次相遇, 保留相遇点的任意一个指针, 将另一个指针设置到head, 重新以相等的速度走, 相遇点就是环入口节点

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
	r.Next = nil
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

//倒数第k个节点(从1开始计数)
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

//找单链表的中点
func middleNode(head *ListNode) *ListNode {
	s := head
	f := head
	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
	}
	return s
}
