package rotatearr

import "fmt"

func rotate2(matrix [][]int) {
	n := len(matrix)
	// Transpose the matrix
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Reverse each row
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for layer := 0; layer < n/2; layer++ {
		first := layer
		last := n - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			// save top
			top := matrix[first][i]
			fmt.Println("top is ", top)
			// left -> top
			matrix[first][i] = matrix[last-offset][first]
			fmt.Println("left is ", matrix[first][i])
			// bottom -> left
			matrix[last-offset][first] = matrix[last][last-offset]
			fmt.Println("bottom is ", matrix[last-offset][first])
			// right -> bottom
			matrix[last][last-offset] = matrix[i][last]
			fmt.Println("right is ", matrix[i][last])
			// top -> right
			matrix[i][last] = top
		}
	}
}
