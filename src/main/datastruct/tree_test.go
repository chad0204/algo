package datastruct

import (
	"fmt"
	"testing"
)

func TestMaxDepth(t *testing.T) {
	root := &TreeNode{1, nil, &TreeNode{2, nil, nil}}
	maxDepth(root)

}

// 104. 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	depth := 0
	traverse(root, &depth)
	return res
}

/*
*
前序遍历, 是一种遍历二叉树的方式, 进阶就是回溯思想
*/
var res int

func traverse(root *TreeNode, depth *int) *int {
	if root == nil {
		return depth
	}
	*depth++
	if root.Left == nil && root.Right == nil {
		res = max(res, *depth)
	}
	traverse(root.Left, depth)
	traverse(root.Right, depth)
	*depth--
	return depth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
	return max(l, r) + 1
}

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
