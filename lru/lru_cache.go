package lru

import "fmt"

type LRU struct {
	capacity int
	head     *Node
	tail     *Node
	cache    map[int]*Node
}

type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

func CreateLRUCache(cap int) *LRU {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.next = head

	return &LRU{capacity: cap, head: head, tail: tail, cache: make(map[int]*Node)}
}

// 1(h),2,4(t) - Cap 3
// get(2)
//2.prev = 2.next
// l.head

func (l *LRU) getValueFromCache(key int) (int, bool) {
	if l.capacity != 0 {
		if node, e := l.cache[key]; e {
			//remove node
			node.prev.next = node.next
			node.next.prev = node.prev
			// add to front
			node.prev = l.head
			node.next = l.head.next
			l.head.next.prev = node
			l.head.next = node
			return node.value, true
		} else {
			return 0, false
		}

	}
	return 0, false
}

// h, 1,2,4, t
// cap -3
// node =2
// h, 2,1,4,t

func (l *LRU) addValueFromCache(key, value int) {
	if node, e := l.cache[key]; e {
		node.value = value
		node.next.prev = node.prev.next
		node.prev.next = node.next
		node.next = l.head.next
		node.prev = l.head
		l.head.next = node
		l.head.next.prev = node
	} else {
		if len(l.cache) >= l.capacity {
			if l.tail.prev == l.head {
				return // List is empty
			}
			oldest := l.tail.prev
			oldest.prev.next = oldest.next
			oldest.next.prev = oldest.prev
			delete(l.cache, oldest.key)
		}
		//add to front
		newNode := &Node{key: key, value: value}

		newNode.prev = l.head
		newNode.next = l.head.next
		l.head.next = newNode
		l.head.next.prev = newNode

		l.cache[key] = newNode
	}
}

func LRUCache() {
	cache := CreateLRUCache(2)
	cache.addValueFromCache(1, 1)
	cache.addValueFromCache(2, 2)
	fmt.Println(cache.getValueFromCache(1)) // returns 1
	cache.addValueFromCache(3, 3)           // evicts key 2
	fmt.Println(cache.getValueFromCache(2)) // returns 0 (not found)
	cache.addValueFromCache(4, 4)           // evicts key 1
	fmt.Println(cache.getValueFromCache(1)) // returns 0 (not found)
	fmt.Println(cache.getValueFromCache(3)) // returns 3
	fmt.Println(cache.getValueFromCache(4)) // retuurn 4
}