package algo

type cumulativeSum2 [][]int

func (c *cumulativeSum2) create(n, m int) {
	*c = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		(*c)[i] = make([]int, m+1)
	}
}

func (c cumulativeSum2) add(x, y, v int) {
	c[x+1][y+1] += v
}

func (c cumulativeSum2) build() {
	n := len(c)
	m := len(c[0])

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			c[i][j] += c[i][j-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			c[j][i] += c[j-1][i]
		}
	}
}

func (c cumulativeSum2) get(sx, sy, tx, ty int) int {
	return c[tx][ty] - c[tx][sy-1] - c[sx-1][ty] + c[sx-1][sy-1]
}
