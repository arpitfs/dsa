package lru

type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

func CreateNewNode(key, value int) *Node {
	return &Node{key: key, value: value}
}
