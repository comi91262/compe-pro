package algo

// DFSによる木DP
func treeDp(n, pre int) {
	var dp [100001]int
	var g [][]int
	dp[n]++

	for _, next := range g[n] {
		if pre == next {
			continue
		}

		treeDp(next, n)
		dp[n] += dp[next]
	}

}

// 部分和DP
func boo() {
	var dp [101][100001]bool
	var n, sum int
	var t []int

	dp[0][0] = true
	for i := 0; i < n; i++ {
		for j := 0; j < sum+1; j++ {
			if j >= t[i] {
				dp[i+1][j] = dp[i][j-t[i]] || dp[i][j]
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}
}

// 区間DP dp[l][r] = min(dp[l][k]+dp[k+1][r], dp[l][r]) or min(dp[l+1][r-1]+v, dp[l][r])
func intervalDp(n int, a []int) {
	inf := 1_000_000_000 + 7
	dp := make([][]int, 2*n)
	for i := 0; i < 2*n; i++ {
		dp[i] = make([]int, 2*n)
	}

	for i := 0; i < 2*n; i++ {
		for j := 0; j < 2*n; j++ {
			dp[i][j] = inf
		}
	}
	for i := 0; i < 2*n-1; i++ {
		dp[i][i+1] = abs(a[i] - a[i+1])
	}

	for interval := 3; interval <= 2*n; interval += 2 {
		for j := 0; j < 2*n-interval; j++ {
			cl, cr := j, j+interval

			dp[cl][cr] = min(dp[cl][cr], dp[cl+1][cr-1]+abs(a[cl]-a[cr]))

			for k := cl; k < cr; k++ {
				dp[cl][cr] = min(dp[cl][cr], dp[cl][k]+dp[k+1][cr])
			}
		}
	}
}
