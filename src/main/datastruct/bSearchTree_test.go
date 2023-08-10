package datastruct

// https://mp.weixin.qq.com/s/kcwz2lyRxxOsC3n11qdVSw
// 96. 不同的二叉搜索树
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

// 95. 不同的二叉搜索树 II
func generateTrees(n int) []*TreeNode {
	return build(1, n)
}

func build(lo, hi int) []*TreeNode {
	res := make([]*TreeNode, 0)
	if lo > hi {
		res = append(res, nil)
		return res
	}
	for i := lo; i <= hi; i++ {
		leftTrees := build(lo, i-1)
		rightTrees := build(i+1, hi)

		//穷举
		for l := range leftTrees {
			for r := range rightTrees {
				left := leftTrees[l]
				right := rightTrees[r]

				root := &TreeNode{i, left, right}
				res = append(res, root)
			}
		}
	}
	return res
}
