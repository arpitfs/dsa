package practice

import "fmt"

type Node struct {
	key, value int
	prev, next *Node
}

type Cache struct {
	cap   int
	cache map[int]*Node
	head  *Node
	tail  *Node
}

func RunCache() {
	cache := CreateCache(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}

func CreateCache(cap int) *Cache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head

	return &Cache{
		cap:   cap,
		cache: make(map[int]*Node),
		head:  head,
		tail:  tail,
	}
}

func (l *Cache) Get(key int) int {
	if node, e := l.cache[key]; e {
		node.prev.next = node.next
		node.next.prev = node.prev

		node.next = l.head.next
		node.prev = l.head
		l.head.next.prev = node
		l.head.next = node
		return node.value
	}
	return -1
}

func (l *Cache) Put(key int, value int) {
	if node, e := l.cache[key]; e {
		node.value = value

		node.prev.next = node.next
		node.next.prev = node.prev

		node.next = l.head.next
		node.prev = l.head
		l.head.next.prev = node
		l.head.next = node
	} else {
		if len(l.cache) >= l.cap {
			if l.tail.prev == l.head {
				return
			}

			node := l.tail.prev
			node.prev.next = node.next
			node.next.prev = node.prev
			delete(l.cache, node.key)
		}

		node := &Node{key: key, value: value}
		l.cache[key] = node
		node.next = l.head.next
		node.prev = l.head
		l.head.next.prev = node
		l.head.next = node
	}
}
