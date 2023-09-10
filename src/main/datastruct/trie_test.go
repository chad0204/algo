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
	trieMap.put("abcd", 1)
	trieMap.put("abc", 2)
	trieMap.put("adc", 2)
	trieMap.put("ab", 3)
	trieMap.put("thea", 5)
	//trieMap.remove("thea")
	//trieMap.remove("ab")

	fmt.Print(trieMap.keysWithPattern("a*c"))

	trie := &Trie{}
	trie.Insert("abc")

}

type Trie struct {
	children [26]*Trie //a ~ z 是97～122正好26个, 用索引下标来存char
	val      int       //0 1 用bool也可以
}

func (this *Trie) Insert(word string) {
	this = InsertHelper(this, word, 0)
}

func InsertHelper(node *Trie, word string, i int) *Trie {
	if node == nil {
		node = &Trie{}
	}
	if i == len(word) {
		node.val = 1
		return node
	}
	c := word[i] - 'a'
	node.children[c] = InsertHelper(node.children[c], word, i+1)
	return node
}

func (this *Trie) Search(word string) bool {
	node := this
	for i := 0; i < len(word); i++ {
		if node == nil {
			return false
		}
		c := word[i] - 'a'
		node = node.children[c]
	}
	if node != nil && node.val != 0 {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for i := 0; i < len(prefix); i++ {
		if node == nil {
			return false
		}
		c := prefix[i] - 'a'
		node = node.children[c]
	}
	return node != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
