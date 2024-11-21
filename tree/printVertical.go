package tree

import "fmt"

func verticalPrint() {
	root := TreeSetUp()
	mapVertical := make(map[int][]int)
	printVerticalTree(root, 0, mapVertical)
}

func printVerticalTree(root *Node, i int, mapVertical map[int][]int) {
	mapTree := traverseTree(root, i, mapVertical)
	for k, val := range mapTree {
		fmt.Println("Key:", k)
		for _, d := range val {
			fmt.Print(d, " ")
		}
	}
}

func traverseTree(root *Node, d int, mapVertical map[int][]int) map[int][]int {
	if root == nil {
		return nil
	}

	mapVertical[d] = append(mapVertical[d], root.data)
	traverseTree(root.left, d-1, mapVertical)
	traverseTree(root.right, d+1, mapVertical)

	return mapVertical
}
