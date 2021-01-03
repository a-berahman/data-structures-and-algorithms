package tree

import (
	"fmt"
	"sync"
)

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
	lock sync.RWMutex
}

func (tree *BinarySearchTree) InsertElement(value int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	var currTree *TreeNode
	currTree = &TreeNode{value, nil, nil}
	if tree.root == nil {
		tree.root = currTree
	} else {
		insertTreeNode(tree.root, currTree)
	}
}

func insertTreeNode(tree *TreeNode, newTree *TreeNode) {

	if tree == nil {
		tree = newTree
		return
	}

	if newTree.value < tree.value {
		if tree.left == nil {
			tree.left = newTree
		} else {

			insertTreeNode(tree.left, newTree)
		}
	} else {
		if tree.right == nil {
			tree.right = newTree
		} else {
			insertTreeNode(tree.right, newTree)
		}
	}

}

// InOrderTraverseTree method
func (tree *BinarySearchTree) InOrderTraverseTree(function func(int)) {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	inOrderTraverseTree(tree.root, function)
}

// inOrderTraverseTree method traverses the left, the root, and the right tree.
func inOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		inOrderTraverseTree(treeNode.left, function)
		function(treeNode.value)
		inOrderTraverseTree(treeNode.right, function)
	}
}

// PreOrderTraverseTree method
func (tree *BinarySearchTree) PreOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	preOrderTraverseTree(tree.root, function)
}

//  preOrderTraverseTree method
func preOrderTraverseTree(treeNode *TreeNode, function func(int)) {

	if treeNode != nil {
		function(treeNode.value)
		preOrderTraverseTree(treeNode.left, function)
		preOrderTraverseTree(treeNode.right, function)
	}
}

// PostOrderTraverseTree method
func (tree *BinarySearchTree) PostOrderTraverseTree(function func(int)) {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	postOrderTraverseTree(tree.root, function)
}

// PostOrderTraverseTree method
func postOrderTraverseTree(treeNode *TreeNode, function func(int)) {

	if treeNode != nil {
		postOrderTraverseTree(treeNode.left, function)
		postOrderTraverseTree(treeNode.right, function)
		function(treeNode.value)
	}
}

//MinNode method
func (tree *BinarySearchTree) MinNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	treeNode := tree.root
	if treeNode == nil {
		return (*int)(nil)
	}
	for {
		if treeNode.left == nil {
			return &treeNode.value
		}
		treeNode = treeNode.left
	}

}

// MaxNode method
func (tree *BinarySearchTree) MaxNode() *int {
	tree.lock.Unlock()
	defer tree.lock.Unlock()

	treenode := tree.root
	if treenode == nil {
		return (*int)(nil)
	}
	for {
		if treenode.right == nil {
			return &treenode.value
		}
		treenode = treenode.right
	}
}

// SearchNode method
func (tree *BinarySearchTree) SearchNode(key int) bool {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	return search(tree.root, key)
}

func search(tree *TreeNode, key int) bool {

	if tree == nil {
		return false
	}
	if key < tree.value {
		return search(tree.left, key)
	}
	if key > tree.value {
		return search(tree.right, key)
	}

	return true

}

func (tree *BinarySearchTree) RemoveNode(key int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	remove(tree.root, key)
}
func remove(treeNode *TreeNode, key int) *TreeNode {
	if treeNode == nil {
		return nil
	}
	if key < treeNode.value {
		treeNode.left = remove(treeNode.left, key)
		return treeNode
	}
	if key > treeNode.value {
		treeNode.right = remove(treeNode.right, key)
		return treeNode
	}
	if treeNode.left == nil {
		treeNode = treeNode.right
		return treeNode
	}
	if treeNode.right == nil {
		treeNode = treeNode.left
		return treeNode
	}

	var leftMostRight *TreeNode
	leftMostRight = treeNode.right
	for {
		if leftMostRight != nil && leftMostRight.left != nil {
			leftMostRight = leftMostRight.left
		} else {
			break
		}
	}
	treeNode.value = leftMostRight.value
	treeNode.right = remove(treeNode.right, treeNode.value)
	return treeNode
}

// String method
func (tree *BinarySearchTree) String() {
	tree.lock.Lock()
	defer tree.lock.Unlock()

	fmt.Println("******************************************")
	stringify(tree.root, 0)
	fmt.Println("******************************************")
}
func stringify(treeNode *TreeNode, level int) {
	if treeNode == nil {
		return
	}

	format := ""
	for i := 0; i < level; i++ {
		format += " "
	}
	format += "---[ "
	level++
	stringify(treeNode.left, level)
	fmt.Printf(format+"%d\n", treeNode.value)
	stringify(treeNode.right, level)
}
