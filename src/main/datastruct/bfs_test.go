package datastruct

import (
	"fmt"
	"testing"
)

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
			if t.Left == nil && t.Right == nil { //没有子节点的节点才算叶子节点
				return level
			}
		}
	}
	return level
}

//二叉树的最小深度 dfs解法
func minDepthDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	v = 999999
	res := 1
	traverseDFS(root, &res)
	return v
}

var v = 999999

func traverseDFS(root *TreeNode, res *int) {
	if root.Left == nil && root.Right == nil {
		v = Min(*res, v)
		return
	}
	*res = *res + 1
	if root.Left != nil {
		traverseDFS(root.Left, res)
	}
	if root.Right != nil {
		traverseDFS(root.Right, res)
	}
	*res = *res - 1
}
