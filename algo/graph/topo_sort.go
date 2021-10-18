package graph

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
