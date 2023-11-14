package sudoku

func isValidSudoku2(board [][]byte) bool {
	var row, col, box [9][9]bool

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := board[i][j] - '1'
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

func isValidSudoku(board [][]byte) bool {
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
	startRow, startCol := row/3*3, col/3*3
	for i := startRow; i < startRow+3; i++ {
		for j := startCol; j < startCol+3; j++ {
			if (i != row || j != col) && board[i][j] == num {
				return false
			}
		}
	}
	return true
}
