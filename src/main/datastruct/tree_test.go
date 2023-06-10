package datastruct

import "testing"

func TestMaxDepth(t *testing.T) {
	root := &TreeNode{1, nil, &TreeNode{2, nil, nil}}
	maxDepth(root)

}

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
