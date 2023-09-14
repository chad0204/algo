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

//迭代
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

//递归
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

//最小前缀, 没有就返回word
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

//递归
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
