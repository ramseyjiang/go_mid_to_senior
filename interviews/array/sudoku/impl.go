package sudoku

func isValidSudoku(board [][]byte) bool {
	// Three 9*9 boolean arrays row, col, and box are initialized.
	// Each array is used to track whether a number (1-9) has already been used in the corresponding row, column, or 3*3 sub-box.
	var row, col, box [9][9]bool

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				// Calculate the value num by subtracting '1' from the byte value of the cell
				// This is to convert the character '1'-'9' to an integer 0-8, suitable for indexing arrays.
				num := board[i][j] - '1'

				// It calculates boxIndex, the index of the 3*3 sub-box to which the cell belongs.
				// This is done by dividing the row and column indices by 3, multiplying the row index by 3, and adding the column index.
				// This ensures that each 3x3 sub-box gets a unique index from 0 to 8.
				boxIndex := (i/3)*3 + j/3

				if row[i][num] || col[j][num] || box[boxIndex][num] {
					return false
				}

				row[i][num], col[j][num], box[boxIndex][num] = true, true, true
			}
		}
	}

	return true
}

func isValidSudoku2(board [][]byte) bool {
	for j := 0; j < 9; j++ {
		for i := 0; i < 9; i++ {
			if board[i][j] != '.' {
				if !validRule(board, i, j) {
					return false
				}
			}
		}
	}
	return true
}

func validRule(board [][]byte, row int, col int) bool {
	num := board[row][col]

	// Check if the number is valid in the row
	for i := 0; i < 9; i++ {
		if i != col && board[row][i] == num {
			return false
		}
	}

	// Check if the number is valid in the column
	for j := 0; j < 9; j++ {
		if j != row && board[j][col] == num {
			return false
		}
	}

	// Check if the number is valid in the 3x3 grid
	startRow, startCol := row/9, col/9
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if (i != row || j != col) && board[i][j] == num {
				return false
			}
		}
	}
	return true
}
