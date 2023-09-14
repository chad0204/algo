package datastruct

// https://mp.weixin.qq.com/s/hGrTUmM1zusPZZ0nA9aaNw

type TrieMap struct {
	root *TrieNode
	size int
}

type TrieNode struct {
	val   int            //初始化0, 所以0表示没有值, TrieMap不能put value == 0的key. 前缀数应该不存值, bool类型即可
	child [256]*TrieNode // 用索引下标来存char, ASCII总共有256位, 包含所有char
}

// 从节点 node 开始搜索 key，如果存在返回对应节点，否则返回 null
func getNode(node *TrieNode, key string) *TrieNode {
	p := node
	for i := 0; i < len(key); i++ {
		if p == nil {
			return nil
		}
		c := key[i]
		p = p.child[c]
	}
	return p
}

func getNodeV2(node *TrieNode, key string, i int) *TrieNode {
	if node == nil {
		return nil
	}
	if len(key) == i {
		return node
	}
	return getNodeV2(node.child[i], key, i+1)
}

func (t *TrieMap) get(key string) int {
	node := getNode(t.root, key)
	//node不为nil或者node的val为空(这里用0表示)， 表示值不存在。存在节点存在val为空是因为下面可能还有值, 比如 存了abc, 没存ab, 找ab
	if node == nil || node.val == 0 {
		return 0
	}
	return node.val
}

// 就算getNode(key)的返回值x非空，也只能说字符串key是一个「前缀」；除非x.val同时非空，才能判断键key存在
func (t *TrieMap) containsKey(key string) bool {
	return t.get(key) != 0
}

func (t *TrieMap) put(key string, val int) {
	if !t.containsKey(key) {
		t.size++
	}
	t.root = t.putHelper(t.root, key, val, 0)
}

func (t *TrieMap) putHelper(node *TrieNode, key string, val int, i int) *TrieNode {
	if node == nil {
		node = &TrieNode{}
	}
	if i == len(key) {
		node.val = val
		return node
	}
	c := key[i]
	node.child[c] = t.putHelper(node.child[c], key, val, i+1)
	return node
}

func (t *TrieMap) remove(key string) {
	if !t.containsKey(key) {
		return
	}
	t.root = removeHelper(t.root, key, 0)
	t.size--
}

func removeHelper(node *TrieNode, key string, i int) *TrieNode {
	if node == nil {
		return nil
	}
	if i != len(key) {
		//递归去子树删除
		c := key[i]
		node.child[c] = removeHelper(node.child[c], key, i+1)
	}
	//后序处理, 倒着向上处理

	if i == len(key) {
		// 找到了 key的最后一个字符对应的 TrieNode，删除 val。 这里不能node = nil, 因为下面可能还有值, 只能把当前节点的值设为空
		node.val = 0
	}

	//路径上存在有值节点, 不能被删除。 比如abcde, 删除e, 发现d有值, d就不能删除。
	if node.val != 0 {
		return node
	}

	//检查是否存在有值子节点（后缀）, 有就不用清理, 没有则清理。比如abcdef, 删除e, 需要判断e下面有没有值, 有f所以节点e不能删除。
	for c := 0; c < 256; c++ {
		if node.child[c] != nil {
			return node
		}
	}
	//既没有存储 val（这里val==0），也没有后缀树枝，上面判断过。比如abcde, 删除e, e的值已经是0, e也没有子值, d = nil
	return nil
}

// 判断是和否存在前缀为 prefix 的键
func (t *TrieMap) hasKeyWithPrefix(prefix string) bool {
	// 只要能找到 prefix 对应的节点，就是存在前缀
	return getNode(t.root, prefix) != nil
}

// 在所有键中寻找 query 的最短前缀 存了abcde, abc， ab。 那么abcdef的最短前缀就是ab
func (t *TrieMap) shortestPrefixOf(query string) string {
	p := t.root
	for i := 0; i < len(query); i++ {
		if p == nil {
			//query = abcd, dic = abd. c对应nil, abd不是前缀
			return ""
		}
		if p.val != 0 {
			//找到最短
			return query[:i]
		}
		c := query[i]
		p = p.child[c]
	}
	if p != nil && p.val != 0 {
		//query本身就是最短前缀
		return query
	}
	return ""
}

// 在所有键中寻找 query 的最长前缀
func (t *TrieMap) longestPrefixOf(query string) string {
	p := t.root
	max := 0
	for i := 0; i < len(query); i++ {
		if p == nil {
			return ""
		}
		if p.val != 0 {
			//更新最大值
			max = i
		}
		c := query[i]
		p = p.child[c]
	}
	if p != nil && p.val != 0 {
		//query本身就是最短前缀
		return query
	}
	return query[:max]
}

// 搜索前缀为 prefix 的所有键
func (t *TrieMap) keysWithPrefix(prefix string) []string {
	res := make([]string, 0)
	//找到prefix对应的节点
	node := getNode(t.root, prefix)
	if node == nil {
		return res
	}
	path := []byte(prefix)

	var traversalTrie func(res *[]string, node *TrieNode, path []byte)
	traversalTrie = func(res *[]string, node *TrieNode, path []byte) {
		if node == nil {
			return
		}
		if node.val != 0 {
			*res = append(*res, string(path))
		}
		for i := 0; i < 256; i++ {
			path = append(path, byte(i))
			traversalTrie(res, node.child[i], path)
			path = path[:len(path)-1]
		}
	}
	traversalTrie(&res, node, path)
	return res
}

// 通配符 * 匹配任意字符
func (t *TrieMap) keysWithPattern(pattern string) []string {
	res := make([]string, 0)
	path := make([]byte, 0)

	var traversalTrie func(res *[]string, path []byte, node *TrieNode, i int, pattern string)
	traversalTrie = func(res *[]string, path []byte, node *TrieNode, i int, pattern string) {
		if node == nil {
			return
		}
		if i == len(pattern) {
			//分支匹配完成
			if node.val != 0 {
				*res = append(*res, string(path))
			}
			return
		}

		c := pattern[i]
		if c == '*' {
			//遍历所有字符
			for j := 0; j < 256; j++ {
				path = append(path, byte(j))
				traversalTrie(res, path, node.child[j], i+1, pattern)
				path = path[:len(path)-1]
			}
		} else {
			//只匹配c的节点
			path = append(path, c)
			traversalTrie(res, path, node.child[c], i+1, pattern)
			path = path[:len(path)-1]
		}
	}
	traversalTrie(&res, path, t.root, 0, pattern)
	return res
}
