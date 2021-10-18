package graph

var used []bool

// DFS 循環グラフ
func dfs(x int, g [][]int, path *[]int) {
	used[x] = true

	for _, next := range g[x] {
		if used[next] {
			continue
		}

		dfs(next, g, path)
	}

	*path = append(*path, x)
}

func scc(n, m int, g, h [][]int) {
	path := make([]int, 0, n)
	used = make([]bool, n+1)
	for i := 1; i < n+1; i++ {
		if used[i] {
			continue
		}
		dfs(i, g, &path)
	}

	used = make([]bool, n+1)
	for i := 1; i < n+1; i++ {
		if used[path[n-i]] {
			continue
		}
		group := make([]int, 0)
		dfs(path[n-i], h, &group)
	}
}
