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
	kv       map[int]*LRUNode // map里的所有元素, 都被维护成链表, map的值是Node指针, 指向链表中的元素
	head     *LRUNode         //head.Next是最旧的元素
	tail     *LRUNode         //tail.Prev是最新的元素
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

/**
插入一个元素, 一定是四步操作, 注意不要先断掉tail和prev的连接
*/
func (this *LRUCache) addLast(node *LRUNode) {
	this.tail.prev.next = node
	node.next = this.tail
	node.prev = this.tail.prev
	this.tail.prev = node
}

func (this *LRUCache) addLast2(node *LRUNode) {
	node.prev = this.tail.prev
	this.tail.prev.next = node
	this.tail.prev = node
	node.next = this.tail
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.kv[key]; !ok {
		return -1
	}
	//移动到队尾
	this.moveLast(this.kv[key])
	return this.kv[key].val
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.kv[key]; ok {
		//已存在, 只需要更新, 移动即可
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
		//删除队头
		this.delNode(this.head.next)
		this.size--
	}
}
