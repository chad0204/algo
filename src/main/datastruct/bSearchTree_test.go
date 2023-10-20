package datastruct

/*

二叉搜索数: 对于节点node, 左子树节点值都比自己小, 右子树节点值都比自己大; 对于节点node, 他的左右子树都是二叉搜索树


隐藏条件: 二叉搜索树的中序遍历是有序的


*/

// 538. 把二叉搜索树转换为累加树
func convertBST(root *TreeNode) *TreeNode {
	res := 0
	convertBSTHelper(root, &res)
	return root
}

func convertBSTHelper(root *TreeNode, res *int) {
	if root == nil {
		return
	}
	//要累加比自己大的元素, 所以从右边开始中序遍历
	convertBSTHelper(root.Right, res)
	*res = *res + root.Val
	root.Val = *res
	convertBSTHelper(root.Left, res)

}

// 230. 二叉搜索树中第K小的元素
func kthSmallest(root *TreeNode, k int) int {
	res := 0
	rank := 0
	kthSmallestHelper(root, k, &res, &rank)
	return res
}

func kthSmallestHelper(root *TreeNode, k int, res *int, rank *int) {
	if root == nil {
		return
	}
	kthSmallestHelper(root.Left, k, res, rank)
	*rank++
	if *rank == k {
		*res = root.Val
		return
	}
	kthSmallestHelper(root.Right, k, res, rank)
}

//98. 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	return isValidHelper(root, nil, nil)
}

func isValidHelper(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	//左子树更新最大值. 右子树更新最小值
	return isValidHelper(root.Left, min, root) && isValidHelper(root.Right, root, max)
}

// 450. 删除二叉搜索树中的节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	return deleteNodeHelper(root, key)
}

func deleteNodeHelper(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < key {
		root.Right = deleteNodeHelper(root.Right, key)
	} else if root.Val > key {
		root.Left = deleteNodeHelper(root.Left, key)
	} else {
		//注意 这里不能直接将右边最小值拿上来，然后拼接上原先的左右。因为最小值还没有被删除。
		//右边的最小值
		minRight := getMin(root.Right)
		if minRight == nil {
			return root.Left
		}
		root.Val = minRight.Val
		root.Right = deleteNodeHelper(root.Right, minRight.Val)
	}
	return root
}

func getMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root
	}
	return getMin(root.Left)
}

// https://mp.weixin.qq.com/s/kcwz2lyRxxOsC3n11qdVSw
// 96. 不同的二叉搜索树
func numTrees(n int) int {
	mem := make([][]int, n+1) //二维数据都是n+1
	for i := range mem {
		mem[i] = make([]int, n+1)
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	return dp(1, n, mem)
}

func dp(lo, hi int, mem [][]int) int {
	if lo > hi {
		//一边倒的情况, 只有一个分支
		return 1
	}
	if mem[lo][hi] != -1 {
		return mem[lo][hi]
	}
	res := 0
	for i := lo; i <= hi; i++ {
		//以i为根
		left := dp(lo, i-1, mem)
		right := dp(i+1, hi, mem)
		res += left * right
	}
	mem[lo][hi] = res
	return res
}

// 95. 不同的二叉搜索树 II
func generateTrees(n int) []*TreeNode {
	return generateTreesHelper(1, n)
}

func generateTreesHelper(min int, max int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if min > max {
		//nil 也是一种节点
		res = append(res, nil)
	}

	//以i逐个作为root
	for rootVal := min; rootVal <= max; rootVal++ {
		lefts := generateTreesHelper(min, rootVal-1)
		rights := generateTreesHelper(rootVal+1, max)
		for _, left := range lefts {
			for _, right := range rights {
				res = append(res, &TreeNode{rootVal, left, right})
			}
		}
	}
	return res
}
