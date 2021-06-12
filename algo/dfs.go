var g [][]int
var a, b []int

func dfs(n, pre int) {
	for _, next := range g[n] {
		if pre == next {
			continue
		}

		dfs(next, n)
	}

}

func input() {
	var n int
	fmt.Fscan(reader, &n)

	g = make([][]int, n+1)
	a = make([]int, n+1)
	b = make([]int, n+1)

	for i := 1; i < n; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
		g[a[i]] = append(g[a[i]], b[i])
		g[b[i]] = append(g[b[i]], a[i])
	}

	dfs(1, 1)
}
