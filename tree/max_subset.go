package tree

import "fmt"

func maxSubset() {
	root := TreeSetUp()
	inc, exc := getMaxSubset(root)
	fmt.Println("Max Subset: ", maximum(inc, exc))

}

func getMaxSubset(root *Node) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftInc, leftExc := getMaxSubset(root.left)
	rightInc, rightExc := getMaxSubset(root.right)

	inc := root.data + leftExc + rightExc
	exc := maximum(leftExc, leftInc) + maximum(rightExc, rightInc)

	return inc, exc
}
