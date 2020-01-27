package problem

import "fmt"

type LinkNode struct {
	key   int
	value int
	pre   *LinkNode
	next  *LinkNode
}
type LRUCache struct {
	capacity int
	head     *LinkNode
	tail     *LinkNode
	hashMap  map[int]*LinkNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{}
	cache.capacity = capacity
	cache.head = &LinkNode{}
	cache.tail = &LinkNode{}
	cache.head.next = cache.tail
	cache.tail.pre = cache.head
	cache.hashMap = make(map[int]*LinkNode, 0)
	return cache
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.hashMap[key]
	if ok {
		if node.pre != this.head {
			nodePre := node.pre
			nodeNext := node.next
			nodePre.next = nodeNext
			nodeNext.pre = nodePre
			headNext := this.head.next
			node.next = headNext
			node.pre = this.head
			this.head.next = node
			fmt.Println(this.head.next)
		}

		return node.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.hashMap[key]
	if ok {
		node.value = value
	} else {
		//删除链表尾部
		if len(this.hashMap) >= this.capacity {
			tailNode := this.tail.pre
			if tailNode != this.head {
				delete(this.hashMap, tailNode.key)
				tailNode.key = key
				tailNode.value = value
				this.hashMap[key] = tailNode
			}
		}
		node := &LinkNode{}
		node.key = key
		node.value = value
		node.pre = this.head
		node.next = this.head.next
		this.head.next.pre = node
		this.head.next = node
		this.hashMap[key] = node
	}
}
