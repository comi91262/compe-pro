package algo

func cumulativeSum2(a [][]int) [][]int {
	m := len(a) + 1

	r := make([][]int, m)
	for i := 0; i < m; i++ {
		r[i] = make([]int, m)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < m; j++ {
			r[i][j] = r[i][j-1] + a[i-1][j-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < m; j++ {
			r[i][j] += r[i-1][j]
		}
	}

	return r
}
