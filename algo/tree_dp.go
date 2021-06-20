package algo

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
