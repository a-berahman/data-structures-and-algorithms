package main

import tree "github.com/a-berahman/data-structures-and-algorithms/Tree"

func main() {

	var tree *tree.BinarySearchTree = &tree.BinarySearchTree{}
	tree.InsertElement(8)
	tree.InsertElement(3)
	tree.InsertElement(10)
	tree.InsertElement(1)
	tree.InsertElement(6)
	tree.String()
}
