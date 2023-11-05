package main

import (
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}
type BST struct {
	root *Node
}

func (bst *BST) Insert(val int) {
	bst.InsertRec(bst.root, val)
}

func (bst *BST) InsertRec(node *Node, val int) *Node {
	if bst.root == nil {
		bst.root = &Node{val, nil, nil}
		return bst.root
	}
	if node == nil {
		return &Node{val, nil, nil}
	}
	if val <= node.data {
		node.left = bst.InsertRec(node.left, val)
	}
	if val > node.data {
		node.right = bst.InsertRec(node.right, val)
	}
	return node
}

func (bst *BST) Search(val int) bool {
	found := bst.SearchRec(bst.root, val)
	return found
}

func (bst *BST) SearchRec(node *Node, val int) bool {
	if node == nil {
		return false
	}

	if node.data == val {
		return true
	}

	if val < node.data {
		return bst.SearchRec(node.left, val)
	}

	return bst.SearchRec(node.right, val)
}

// Inorder traversal is a depth-first, recursive method that traverses the tree in a left node > root node > right node order.
// A depth-first approach will poke the deepest leaf nodes in a subtree before moving on to the next subtree.
func (bst *BST) Inorder(node *Node) {
	if node == nil {
		return
	}

	bst.Inorder(node.left)
	bst.Inorder(node.right)
	fmt.Print(node.data, " ")
}

// LevelOrder traversal takes a breadth-first approach.
// A breadth-first approach, unlike a depth-first approach, will iterate through the tree level-by-level.
func (bst *BST) LevelOrder() {
	if bst.root == nil {
		return
	}
	nodeList := make([]*Node, 0)
	nodeList = append(nodeList, bst.root)
	for len(nodeList) != 0 {
		current := nodeList[0]
		fmt.Print(current.data, " ")
		if current.left != nil {
			nodeList = append(nodeList, current.left)
		}
		if current.right != nil {
			nodeList = append(nodeList, current.right)
		}
		nodeList = nodeList[1:]
	}
}

// In BST, it has two most popular methods: inorder traversal and level order traversal.
// A binary tree is a tree where each node can have at most two children.
// A binary search tree is a special type of binary tree.
// The left child must have a value less than the parent, and the right child must have a value greater than the parent.
// Binary search trees, abbreviated as BSTs.
// Perform a binary search: split the list in half, pick the half where the item would be, split again, and vice versa.
// While a linear search would take O(n) time, a binary search would take O(log n) time, which makes it more efficient.
func main() {
	bst := BST{}
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(15)
	bst.Insert(20)
	bst.Insert(17)
	bst.Insert(4)
	bst.Insert(6)
	bst.Inorder(bst.root)
	fmt.Println()
	bst.LevelOrder()
	fmt.Println()
	fmt.Println(bst.Search(5))
	fmt.Println(bst.Search(11))
}
