package algo

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

// トポロジカルソートを行う関数
func topoSort(g [][]int) []int {
	var ans []int

	n := len(g)

	index := make([]int, n)
	for i := 0; i < n; i++ {
		for _, e := range g[i] {
			index[e]++
		}
	}

	var que = make(queue, 0)
	for i := 0; i < n; i++ {
		if index[i] == 0 {
			que.push(i)
		}
	}

	for !que.empty() {
		now := que.pop()
		ans = append(ans, now)
		for _, e := range g[now] {
			index[e]--
			if index[e] == 0 {
				que.push(e)
			}
		}
	}
	return ans
}
