package tree

import "fmt"

func diameter() {
	root := TreeSetUp()
	diameter := getDiameter(root)
	fmt.Println("Diamter: ", diameter)
}

func getDiameter(root *Node) int {
	if root == nil {
		return 0
	}

	d1 := height(root.left) + height(root.right)
	d2 := getDiameter(root.left)
	d3 := getDiameter(root.right)

	return maximum(d1, maximum(d2, d3))

}

func height(root *Node) int {
	if root == nil {
		return 0
	}

	h1 := height(root.left)
	h2 := height(root.right)

	return maximum(h1, h2) + 1
}

func maximum(h1, h2 int) int {
	if h1 > h2 {
		return h1
	}
	return h2
}
