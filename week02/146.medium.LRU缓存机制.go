package week02

// https://leetcode-cn.com/problems/lru-cache/
// https://leetcode-cn.com/submissions/detail/183321023/

type LRUCache struct {
	capacity   int
	dict       map[int]*Node
	head, tail *Node
}

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		dict:     make(map[int]*Node, capacity),
		head:     &Node{},
		tail:     &Node{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) removeDLinkListNode(n *Node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (this *LRUCache) addNodeToDLinkListTail(n *Node) {
	n.prev = this.tail.prev
	n.next = this.tail
	this.tail.prev.next = n
	this.tail.prev = n
}

/*
	if !存在
		return -1
	else
		更新节点到表尾
		返回值
*/

func (this *LRUCache) Get(key int) int {
	n, ok := this.dict[key]
	if !ok {
		return -1
	}
	this.removeDLinkListNode(n)
	this.addNodeToDLinkListTail(n)
	return n.value
}

/*
	if 存在
		更新节点到表尾
	else
		if 满了
			淘汰表头
		新建node
		插入节点到表尾
*/

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.dict[key]; ok {
		n.value = value
		this.removeDLinkListNode(n)
		this.addNodeToDLinkListTail(n)
		return
	}
	if len(this.dict) == this.capacity {
		evicted := this.head.next
		delete(this.dict, evicted.key)
		this.removeDLinkListNode(evicted)
	}
	n := &Node{
		key:   key,
		value: value,
	}
	this.dict[key] = n
	this.addNodeToDLinkListTail(n)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
