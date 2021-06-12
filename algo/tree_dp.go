var dp [100001]int
var g [][]int

func dfs(n, pre int) {
	dp[n]++

	for _, next := range g[n] {
		if pre == next {
			continue
		}

		dfs(next, n)
		dp[n] += dp[next]
	}

}
