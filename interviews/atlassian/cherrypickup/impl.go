package cherrypickup

func cherryPickup2(grid [][]int) int {
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

func cherryPickup(grid [][]int) int {
	n, dp, inf := len(grid), make([][]int, len(grid)), -1<<31-1
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = inf
		}
	}
	dp[0][0] = grid[0][0]
	minTimes := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	maxTimes := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for k := 1; k < n*2-1; k++ {
		for x1 := minTimes(k, n-1); x1 >= maxTimes(k-n+1, 0); x1-- {
			for x2 := minTimes(k, n-1); x2 >= x1; x2-- {
				y1, y2 := k-x1, k-x2
				if grid[x1][y1] == -1 || grid[x2][y2] == -1 {
					dp[x1][x2] = inf
					continue
				}
				res := dp[x1][x2] // 都往右
				if x1 > 0 {
					res = maxTimes(res, dp[x1-1][x2]) // 往下，往右
				}
				if x2 > 0 {
					res = maxTimes(res, dp[x1][x2-1]) // 往右，往下
				}
				if x1 > 0 && x2 > 0 {
					res = maxTimes(res, dp[x1-1][x2-1]) // 都往下
				}
				res += grid[x1][y1]
				if x2 != x1 { // 避免重复摘同一个樱桃
					res += grid[x2][y2]
				}
				dp[x1][x2] = res
			}
		}
	}
	return maxTimes(dp[n-1][n-1], 0)
}
