package datastruct

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

// https://mp.weixin.qq.com/s/kcwz2lyRxxOsC3n11qdVSw
// 96. 不同的二叉搜索树
// 95. 不同的二叉搜索树 II
func numTrees(n int) int {
	mem := make([][]int, n+1)
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
