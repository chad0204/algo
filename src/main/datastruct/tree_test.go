package datastruct

import (
	"fmt"
	"testing"
)

/**
遍历的模式, 回溯思想, 前序遍历



分解子问题的模式, 动态规划。分解子问题一般需要拿到子问题的结果, 也就是需要返回值



*/
func TestTreeNode_PreTraverse(t *testing.T) {
	root := &TreeNode{0,
		&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}},
		&TreeNode{4, &TreeNode{5, nil, nil}, &TreeNode{6, nil, nil}}}
	var pres []int
	pres = root.PreTraverse(pres)
	fmt.Println(pres)
}

func TestTreeNode_PostTraverse(t *testing.T) {
	root := &TreeNode{6,
		&TreeNode{2, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}},
		&TreeNode{5, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}}
	var posts []int
	root.PostTraverse(&posts)
	fmt.Println(posts)
}

func TestTreeNode_InTraverse(t *testing.T) {
	root := &TreeNode{3,
		&TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{2, nil, nil}},
		&TreeNode{5, &TreeNode{4, nil, nil}, &TreeNode{6, nil, nil}}}
	var inorders []int
	root.InTraverse(&inorders)
	fmt.Println(inorders)
}

func TestTreeNode_LevelTraverse(t *testing.T) {
	root := &TreeNode{0,
		&TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{2, &TreeNode{5, nil, nil}, &TreeNode{6, nil, nil}}}
	var levels []int
	root.LevelTraverse(&levels)
	fmt.Println(levels)
}

func TestMaxDepth(t *testing.T) {
	root := &TreeNode{1, nil, &TreeNode{2, nil, nil}}
	maxDepth(root)

}

// 104. 二叉树的最大深度 涉及子树, 需要后续遍历并设置返回值 左右根
func maxDepth(root *TreeNode) int {
	res = 0 //每次执行 清空
	depth := 0
	traverse(root, &depth)
	return res
}

/*
*
前序遍历, 是一种遍历二叉树的方式, 进阶就是回溯思想
*/
var res int

func traverse(root *TreeNode, depth *int) {
	if root == nil {
		return
	}
	*depth++
	if root.Left == nil && root.Right == nil {
		res = Max(res, *depth)
	}
	traverse(root.Left, depth)
	traverse(root.Right, depth)
	*depth--
}

/*
后序遍历, 是一种分解子问题的思想, 进阶就是动态规划
*/
func traverseV2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := traverseV2(root.Left)
	r := traverseV2(root.Right)
	return Max(l, r) + 1
}

// 543. 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	/*
		思路是"所有"子树的左右最大深度之和。
		涉及左右子树所以是后序遍历, "最大"需要记录最大值
	*/
	d := 0
	diameter(root, &d)
	return d
}

func diameter(root *TreeNode, d *int) int {
	if root == nil {
		return 0
	}
	l := diameter(root.Left, d)
	r := diameter(root.Right, d)
	*d = Max(l+r, *d)
	return Max(l, r) + 1 // 计算左右子树的深度
}

// 114. 二叉树展开为链表, 分解子问题模式。如果返回类型不是void, 可以考虑遍历的思想
func flatten(root *TreeNode) {
	traversal(root)
}

/**

0. 开始
  1
 / \
2   3

1. 记录 l, r
2 , 3

2. 左接到右边
  1
 / \
    2

3. 再把老的右接到新的右边
  1
 / \
	2
	 \
      3

*/
func traversal(root *TreeNode) {
	if root == nil {
		return
	}

	traversal(root.Left)
	traversal(root.Right)

	//记录左右子树
	l := root.Left
	r := root.Right

	//左为空, 右为左
	root.Left = nil
	root.Right = l

	//最后处理r, 接到新右的最后
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = r
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 116. 填充每个节点的下一个右侧节点指针. 解决子树直接的空隙问题
//解法一: 层序遍历
func connectLevel(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		var temp *Node
		for i := 0; i < n; i++ {
			curr := queue[0]
			queue = queue[1:]
			if curr.Left == nil || curr.Right == nil {
				continue
			}
			queue = append(queue, curr.Left)
			queue = append(queue, curr.Right)
			if temp != nil {
				temp.Next = curr.Left
			}
			curr.Left.Next = curr.Right
			temp = curr.Right
		}
	}
	return root
}

//解法二: 三叉树
func connect3Tree(root *Node) *Node {
	if root == nil {
		return nil
	}
	traversal3Tree(root.Left, root.Right)
	return root
}

func traversal3Tree(node1 *Node, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}
	node1.Next = node2
	traversal3Tree(node1.Left, node1.Right)
	traversal3Tree(node1.Right, node2.Left)
	traversal3Tree(node2.Left, node2.Right)
}

// 106. 从中序与后序遍历序列构造二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	return buildTreeTraversal(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

func buildTreeTraversal(inorder []int, inStart int, inEnd int, postorder []int, postStart int, postEnd int) *TreeNode {
	if postStart > postEnd { //post in 都ok
		return nil
	}
	rootVal := postorder[postEnd]
	//找到分割左右子树的索引, 并计算左右子树长度
	rootIndex := 0
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == rootVal {
			rootIndex = i
			break
		}
	}
	leftSize := rootIndex - inStart
	left := buildTreeTraversal(inorder, inStart, rootIndex-1, postorder, postStart, postStart+leftSize-1)
	right := buildTreeTraversal(inorder, rootIndex+1, inEnd, postorder, postStart+leftSize, postEnd-1)
	return &TreeNode{rootVal, left, right}
}

func TestConstructFromPrePost(t *testing.T) {
	//root := &TreeNode{1,
	//	&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}},
	//	&TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}}
	constructFromPrePost([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 5, 2, 6, 3, 7, 1})

}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	return construct(preorder, 0, len(preorder)-1, postorder, 0, len(postorder)-1)
}

func construct(preorder []int, preStart int, preEnd int,
	postorder []int, postStart int, postEnd int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	//因为需要找个左根, 所以相等要特殊处理
	if preStart == preEnd {
		return &TreeNode{preStart, nil, nil}
	}

	rootVal := preorder[preStart]
	//假设前序遍历的第二个节点是左根
	leftRootVal := preorder[preStart+1]
	//找到左根在后序遍历中的位置
	leftRootIndex := 0
	for i := postStart; i <= postEnd; i++ {
		if postorder[i] == leftRootVal {
			leftRootIndex = i
			break
		}
	}
	leftSize := leftRootIndex - postStart + 1
	left := construct(preorder, preStart+1, preStart+leftSize, postorder, postStart, leftRootIndex)
	right := construct(preorder, preStart+leftSize+1, preEnd, postorder, leftRootIndex+1, postEnd-1)
	return &TreeNode{rootVal, left, right}
}
