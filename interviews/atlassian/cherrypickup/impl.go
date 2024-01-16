package cherrypickup

func cherryPickup(grid [][]int) int {
	n := len(grid)
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	var solve func(int, int, int) int
	solve = func(r1, c1, c2 int) int {
		r2 := r1 + c1 - c2

		// Base cases
		if r1 >= n || c1 >= n || r2 >= n || c2 >= n || grid[r1][c1] == -1 || grid[r2][c2] == -1 {
			return -1e8
		}

		if r1 == n-1 && c1 == n-1 {
			return grid[r1][c1]
		}

		if dp[r1][c1][c2] != -1 {
			return dp[r1][c1][c2]
		}

		cnt := grid[r1][c1]
		if r1 != r2 {
			cnt += grid[r2][c2]
		}

		a := solve(r1+1, c1, c2)
		b := solve(r1+1, c1, c2+1)
		c := solve(r1, c1+1, c2)
		d := solve(r1, c1+1, c2+1)

		dp[r1][c1][c2] = cnt + max(max(a, b), max(c, d))
		return dp[r1][c1][c2]
	}

	result := max(0, solve(0, 0, 0))
	return result
}
