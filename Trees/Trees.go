package main

type treeNode struct {
	value int
	left  *treeNode
	right *treeNode
}

func insert(root *treeNode, value int) {
	if root == nil {
		root = &treeNode{value: value}
	} else if value < root.value {
		insert(root.left, value)
	} else {
		insert(root.right, value)
	}
}

func display(root *treeNode) {
	if root != nil {
		display(root.left)
		print(root.value, " ")
		display(root.right)
	}
}

func main() {
	var root *treeNode
	insert(root, 1)
	insert(root, 2)
	insert(root, 3)
	display(root)
}
