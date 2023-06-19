package datastruct

import (
	"fmt"
	"testing"
)

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

// 104. 二叉树的最大深度 涉及子树, 需要后续遍历并设置返回值 左右根
func TestMaxDepth(t *testing.T) {
	root := &TreeNode{1, nil, &TreeNode{2, nil, nil}}
	maxDepth(root)

}

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

func TestMinDepth(t *testing.T) {
	root := &TreeNode{0,
		&TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}},
		&TreeNode{2, nil, nil}}

	//root := &TreeNode{1,
	//	nil,
	//	&TreeNode{2,
	//		nil,
	//		&TreeNode{3,
	//			nil,
	//			&TreeNode{4,
	//				nil,
	//				&TreeNode{5, nil ,nil}}}}}

	fmt.Println(minDepth(root))
}

// 111. 二叉树的最小深度, 最小, 层序遍历 找到第一个叶子节点的分支即可, 如果用dfs(参考最大深度), 需要遍历所有的分支, 时间复杂度划不来
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var queue []*TreeNode
	queue = append(queue, root)
	level := 0

	for len(queue) != 0 {
		level++
		l := len(queue) //记录长度快照, queue的长度是动态的
		for i := 0; i < l; i++ {
			t := queue[0]
			queue = queue[1:] //需要缩减长度
			if t.Left != nil {
				queue = append(queue, t.Left)
			}
			if t.Right != nil {
				queue = append(queue, t.Right)
			}
			if t.Left == nil && t.Right == nil {
				return level
			}
		}
	}
	return level
}
