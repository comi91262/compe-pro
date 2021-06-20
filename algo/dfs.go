package algo

func dfs(n, pre int, g *[][]int) {
	for _, next := range (*g)[n] {
		if pre == next {
			continue
		}

		dfs(next, n, g)
	}
}
