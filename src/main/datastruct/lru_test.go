package datastruct

import "testing"

func TestLRU(t *testing.T) {
	//["LRUCache","put","put","get","put","get","put","get","get","get"]
	//[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	//[[2],[1,0],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]
	lru := ConstructorLRU(2)
	lru.Put(1, 0)
	lru.Put(2, 2)
	lru.Get(1)
	lru.Put(3, 3)
	lru.Get(2)
	lru.Put(4, 4)
	lru.Get(1)
	lru.Get(3)
	lru.Get(4)

}

type LRUCache struct {
	kv       map[int]*LRUNode
	head     *LRUNode
	tail     *LRUNode
	size     int
	capacity int
}

type LRUNode struct {
	key  int
	val  int
	next *LRUNode
	prev *LRUNode
}

func ConstructorLRU(capacity int) LRUCache {
	head := &LRUNode{-1, -1, nil, nil}
	tail := &LRUNode{-1, -1, nil, nil}
	head.next = tail
	tail.prev = head
	return LRUCache{make(map[int]*LRUNode), head, tail, 0, capacity}
}

func (this *LRUCache) moveLast(node *LRUNode) {
	//原地删除
	this.delNode(node)
	//添加到队尾
	this.addLast(node)
}

func (this *LRUCache) delNode(node *LRUNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = nil
}

func (this *LRUCache) addLast(node *LRUNode) {
	this.tail.prev.next = node
	node.next = this.tail
	node.prev = this.tail.prev
	this.tail.prev = node
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.kv[key]; !ok {
		return -1
	}
	this.moveLast(this.kv[key])
	return this.kv[key].val
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.kv[key]; ok {
		this.kv[key].val = value
		this.moveLast(this.kv[key])
		return
	}
	node := &LRUNode{key, value, nil, nil}
	this.addLast(node)
	this.kv[key] = node
	this.size++
	if this.size > this.capacity {
		delete(this.kv, this.head.next.key)
		this.delNode(this.head.next)
		this.size--
	}
}
