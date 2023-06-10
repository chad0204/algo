package datastruct

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PreTraverse 根左右 从根节点就能得到需要的参数, 解题时一般不用返回值
func (t *TreeNode) PreTraverse(res []int) []int {
	if t == nil {
		return res
	}
	res = append(res, t.Val)
	res = t.Left.PreTraverse(res)
	res = t.Right.PreTraverse(res)
	return res
}

// PostTraverse 左右根 需要遍历完左右子树才能得到需要的参数, 算法大概率要设置返回值
func (t *TreeNode) PostTraverse(res *[]int) {
	if t == nil {
		return
	}
	t.Left.PostTraverse(res)
	t.Right.PostTraverse(res)
	*res = append(*res, t.Val)
}

// InTraverse 左根右 快排 整棵树从左到右有序 BST(二叉搜索树)
func (t *TreeNode) InTraverse(res *[]int) {
	if t == nil {
		return
	}
	t.Left.InTraverse(res)
	*res = append(*res, t.Val)
	t.Right.InTraverse(res)
}
