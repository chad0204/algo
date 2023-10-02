package datastruct

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

/*
*
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
		思路是"所有"子树的左右最大深度之和。注意是所有, 而不是root左右之和
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
	// 返回当前节点的最长子树
	return Max(l, r) + 1
}

// 114. 二叉树展开为链表, 分解子问题模式。如果返回类型不是void, 可以考虑遍历的思想
func flatten(root *TreeNode) {
	traversal(root)
}

/*
*

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

func TestCL(t *testing.T) {
	root := &Node{1,
		&Node{2, &Node{4, nil, nil, nil}, &Node{5, nil, nil, nil}, nil},
		&Node{3, &Node{6, nil, nil, nil}, &Node{7, nil, nil, nil}, nil},
		nil}
	connectLevel(root)
}

// 116. 填充每个节点的下一个右侧节点指针. 解决子树直接的空隙问题
// 解法一: 层序遍历
func connectLevel(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		//每一层开始前清空
		var prevRight *Node
		for i := 0; i < n; i++ {
			curr := queue[0]
			queue = queue[1:]
			if curr.Left == nil || curr.Right == nil {
				continue
			}
			queue = append(queue, curr.Left)
			queue = append(queue, curr.Right)
			curr.Left.Next = curr.Right
			if prevRight != nil {
				prevRight.Next = curr.Left
			}
			//记录当前节点的右节点
			prevRight = curr.Right
		}
	}
	return root
}

// 解法二: 三叉树
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
		return &TreeNode{preorder[preStart], nil, nil}
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

/*
652. 寻找重复的子树
https://mp.weixin.qq.com/s/LJbpo49qppIeRs-FbgjsSQ
*/
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	valMap := make(map[string]int)
	res := make([]*TreeNode, 0)
	traversalFDS(root, valMap, &res)
	return res
}

func traversalFDS(root *TreeNode, valMap map[string]int, res *[]*TreeNode) string {
	if root == nil {
		return "#"
	}

	left := traversalFDS(root.Left, valMap, res)
	right := traversalFDS(root.Right, valMap, res)

	subTree := left + ", " + right + ", " + strconv.Itoa(root.Val)

	valMap[subTree] = valMap[subTree] + 1
	if valMap[subTree] == 2 {
		*res = append(*res, root)
	}
	return subTree
}

// 297. 二叉树的序列化与反序列化
// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247485871&idx=1&sn=bcb24ea8927995b585629a8b9caeed01&chksm=9bd7f7a7aca07eb1b4c330382a4e0b916ef5a82ca48db28908ab16563e28a376b5ca6805bec2&scene=21#wechat_redirect
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	left := this.serialize(root.Left)
	right := this.serialize(root.Right)
	return left + "," + right + "," + strconv.Itoa(root.Val)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	postorder := strings.Split(data, ",")
	return this.deserializeHelper(&postorder)
}

// # # 2     #  #  4  # # 5  3     1
func (this *Codec) deserializeHelper(postorder *[]string) *TreeNode {
	// if len(*postorder) == 0 {
	//     return nil
	// }

	lastIndex := len(*postorder) - 1
	val := (*postorder)[lastIndex]
	//主要就是postorder每次构建完一个节点 都要减掉一个
	*postorder = (*postorder)[:lastIndex]
	if val == "#" {
		return nil
	}

	right := this.deserializeHelper(postorder)
	left := this.deserializeHelper(postorder)

	v, _ := strconv.Atoi(val)
	return &TreeNode{v, left, right}

}

func TestCodec(t *testing.T) {
	root := &TreeNode{1,
		&TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}},
		&TreeNode{3, nil, nil}}
	ser := Constructor()
	deser := Constructor()
	data := ser.serialize(root)
	fmt.Println(data)
	ans := deser.deserialize(data)
	fmt.Println(ans)
}

// 236. 二叉树的最近公共祖先
// https://mp.weixin.qq.com/s/njl6nuid0aalZdH5tuDpqQ
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return lcaHelper(root, p, q)
}

func lcaHelper(root *TreeNode, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val {
		return root
	}
	if root.Val == q.Val {
		return root
	}

	l := lcaHelper(root.Left, p, q)
	r := lcaHelper(root.Right, p, q)

	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	//左右子树都不为空, 找到lca
	return root
}

// 235. 二叉搜索树的最近公共祖先
func lowestCommonAncestor235(root, p, q *TreeNode) *TreeNode {
	if p.Val < q.Val {
		return lcaHelper(root, p, q)
	} else {
		return lcaHelper(root, q, p)
	}
}

func lcaHelper235(root, min, max *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// root min max
	if root.Val > max.Val {
		return lcaHelper235(root.Left, min, max)
	}
	// min max root
	if root.Val < min.Val {
		return lcaHelper235(root.Right, min, max)
	}
	// min root max 找到lca
	return root
}
