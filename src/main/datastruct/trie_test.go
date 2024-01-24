package datastruct

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {

	for i := 0; i < 26; i++ {
		fmt.Printf("i := %v, i + 'a' = %v, char = %v, char = %v \n", i, i+'a', string(byte(i+'a')), string(byte(i)))
	}

	trieMap := &TrieMap{}
	trieMap.put("a", 1)
	trieMap.put("b", 2)
	trieMap.put("c", 2)
	trieMap.put("d", 3)
	trieMap.put("e", 5)
	trieMap.get("f")
	trieMap.put("abc", 6)
	trieMap.put("ab", 100)
	trieMap.remove("thea")
	trieMap.remove("abc")
	trieMap.remove("a")
	trieMap.remove("b")
	fmt.Println(trieMap.get("abc"))
	fmt.Println(trieMap.get("a"))
	fmt.Println(trieMap.get("ab"))

	fmt.Print(trieMap.keysWithPattern("a*c"))

	trie := &Trie{}
	trie.Insert("ab")
	fmt.Println(trie.shortestPrefixV2("abc"))
}

type Trie struct {
	children [26]*Trie //a ~ z 是97～122正好26个, 用索引下标来存char
	val      bool      //0 1 用bool也可以
	Word     string    //0 1 用bool也可以
}

func (this *Trie) Insert(word string) {
	InsertHelper(this, word, 0)
}

func InsertHelper(node *Trie, word string, i int) *Trie {
	if node == nil {
		node = &Trie{}
	}
	if i == len(word) {
		node.val = true
		return node
	}
	c := word[i] - 'a'
	node.children[c] = InsertHelper(node.children[c], word, i+1)
	return node
}

// 迭代
func (this *Trie) Search(word string) bool {
	node := this
	for i := 0; i < len(word); i++ {
		if node == nil {
			return false
		}
		c := word[i] - 'a'
		node = node.children[c]
	}
	if node != nil && node.val {
		return true
	}
	return false
}

// 递归
func (this *Trie) SearchV2(word string) bool {
	node := searchHelper(word, 0, this)
	if node != nil && node.val {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	return searchHelper(prefix, 0, this) != nil
}

func searchHelper(word string, i int, node *Trie) *Trie {
	if node == nil {
		return nil
	}
	if i == len(word) {
		return node
	}
	c := word[i] - 'a'
	return searchHelper(word, i+1, node.children[c])
}

// 最小前缀, 没有就返回word
func (t *Trie) shortestPrefix(word string) string {
	p := t
	for i := 0; i < len(word); i++ {
		if p == nil {
			return word
		}
		if p.val {
			return word[:i]
		}
		c := word[i] - 'a'
		p = p.children[c]
	}
	if p != nil && p.val {
		return word
	}
	return word
}

// 递归
func (t *Trie) shortestPrefixV2(word string) string {
	return shortestPrefixHelper(t, word, 0)
}

func shortestPrefixHelper(node *Trie, word string, i int) string {
	if node == nil {
		return word
	}
	if node.val {
		return word[:i]
	}
	if i == len(word) {
		return word
	}
	c := word[i] - 'a'
	return shortestPrefixHelper(node.children[c], word, i+1)
}

// 212. 单词搜索 II
func findWords(board [][]byte, words []string) []string {
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	trie := &Trie{}
	for _, word := range words {
		addNode(trie, word, 0)
	}

	res := make(map[string]bool)
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			dfsFindWords(board, visited, r, c, trie, res)
		}
	}

	result := make([]string, 0)
	for k := range res {
		result = append(result, k)
	}
	return result
}

func dfsFindWords(board [][]byte, visited [][]bool, r, c int, trie *Trie, res map[string]bool) {
	if r < 0 || c < 0 || r >= len(board) || c >= len(board[0]) {
		return
	}
	//不能漏掉头节点
	ch := board[r][c] - 'a'
	trie = trie.children[ch]
	if visited[r][c] {
		return
	}

	if trie == nil {
		//说明没有这个前缀, 剪枝
		return
	}

	if len(trie.Word) > 0 {
		//说明找到了, 这里不能return。 比如ab, abb同时存在, 匹配到ab, 还要匹配abb
		res[trie.Word] = true
		//return
	}

	visited[r][c] = true
	dfsFindWords(board, visited, r-1, c, trie, res)
	dfsFindWords(board, visited, r+1, c, trie, res)
	dfsFindWords(board, visited, r, c-1, trie, res)
	dfsFindWords(board, visited, r, c+1, trie, res)
	visited[r][c] = false
}

func addNode(root *Trie, word string, i int) *Trie {
	if root == nil {
		root = &Trie{}
	}
	if i == len(word) {
		root.Word = word
		return root
	}
	ch := word[i] - 'a'
	root.children[ch] = addNode(root.children[ch], word, i+1)
	return root
}
