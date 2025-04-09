package round2

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

func (l *Cache) Get(k int) int {
	if node, e := l.cache[k]; e {
		node.next.prev = node.prev
		node.prev.next = node.next

		l.head.next = node
		node.prev = l.head
		l.head.next.prev = node
		node.next = l.head.next

		return node.value

	}
	return -1
}

func (l *Cache) Put(k, v int) {
	if node, e := l.cache[k]; e {
		node.value = v
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

		newNode := &Node{key: k, value: v}
		l.cache[k] = newNode
		newNode.next = l.head.next
		l.head.next.prev = newNode
		l.head.next = newNode
		newNode.prev = l.head
	}
}

func runCache() {
	cache := CreateCache(0)

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
