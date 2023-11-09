package leafsimilar

import "testing"

// Helper function to construct a binary tree from a level-order list (nil represented by -1)
func constructTree(levelOrder []int) *TreeNode {
	if len(levelOrder) == 0 {
		return nil
	}

	root := &TreeNode{Val: levelOrder[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if i < len(levelOrder) && levelOrder[i] != -1 {
			node.Left = &TreeNode{Val: levelOrder[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(levelOrder) && levelOrder[i] != -1 {
			node.Right = &TreeNode{Val: levelOrder[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func TestLeafSimilar(t *testing.T) {
	tests := []struct {
		name  string
		tree1 []int
		tree2 []int
		want  bool
	}{
		{
			name:  "both trees empty",
			tree1: []int{},
			tree2: []int{},
			want:  true,
		},
		{
			name:  "trees with same leaves",
			tree1: []int{3, 5, 1, 6, 2, 9, 8, -1, -1, 7, 4},
			tree2: []int{3, 5, 1, 6, 7, 4, 2, -1, -1, -1, -1, -1, -1, 9, 8},
			want:  true,
		},
		{
			name:  "trees with different leaves",
			tree1: []int{1, 2, 3},
			tree2: []int{1, 3, 2},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root1 := constructTree(tt.tree1)
			root2 := constructTree(tt.tree2)

			leaves1 := getLeaves(root1)
			leaves2 := getLeaves(root2)

			// Debugging: Print the leaves to see what is being compared
			t.Log("Leaves of tree1:", leaves1)
			t.Log("Leaves of tree2:", leaves2)

			if got := leafSimilar(root1, root2); got != tt.want {
				t.Errorf("leafSimilar() = %v, want %v", got, tt.want)
			}
		})
	}
}
