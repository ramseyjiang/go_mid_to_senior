package leafsimilar

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leafSimilar checks if two binary trees are leaf-similar.
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1 := getLeaves(root1)
	leaves2 := getLeaves(root2)

	if len(leaves1) != len(leaves2) {
		return false
	}

	for i := range leaves1 {
		if leaves1[i] != leaves2[i] {
			return false
		}
	}

	return true
}

// getLeaves traverses the tree and returns the leaf value sequence.
func getLeaves(root *TreeNode) []int {
	var leaves []int
	var stack []*TreeNode
	currentNode := root

	for currentNode != nil || len(stack) > 0 {
		for currentNode != nil {
			stack = append(stack, currentNode)
			currentNode = currentNode.Left
		}
		currentNode = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if currentNode.Left == nil && currentNode.Right == nil {
			leaves = append(leaves, currentNode.Val)
		}

		currentNode = currentNode.Right
	}

	return leaves
}
