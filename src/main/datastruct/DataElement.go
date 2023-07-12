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

// LevelTraverse 层序遍历, 就是迭代思想 比较简单
func (t *TreeNode) LevelTraverse(res *[]int) {
	if t == nil {
		return
	}
	levels := make([]*TreeNode, 0)
	levels = append(levels, t)
	for len(levels) > 0 {
		l := len(levels)
		for i := 0; i < l; i++ { //这里的循环是进行分层, 如果不需要记录层级, 其实也可以去掉
			curr := levels[0]
			levels = levels[1:] //相当于pop一个
			*res = append(*res, curr.Val)
			if curr.Left != nil {
				levels = append(levels, curr.Left)
			}
			if curr.Right != nil {
				levels = append(levels, curr.Right)
			}
		}
	}
}

func (t *TreeNode) LevelTraverseV2(res *[]int) {
	if t == nil {
		return
	}
	queue := make([]*TreeNode, 1)
	queue[0] = t
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		*res = append(*res, curr.Val)
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
