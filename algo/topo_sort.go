package algo

type queue []int

func (q *queue) push(n int) {
	*q = append(*q, n)
}

func (q *queue) pop() int {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

func (q *queue) empty() bool {
	return len(*q) == 0
}

func topoSort(g [][]int) []int { // bfs
	var ans []int

	n := len(g)
	ind := make([]int, n)
	for i := 0; i < n; i++ { // 次数を数えておく
		for e := range g[i] {
			ind[e]++
		}
	}
	var que = make(queue, 0)
	for i := 0; i < n; i++ { // 次数が0の点をキューに入れる
		if ind[i] == 0 {
			que.push(i)
		}
	}
	for !que.empty() {
		now := que.pop()
		ans = append(ans, now)
		for e := range g[now] {
			ind[e]--
			if ind[e] == 0 {
				que.push(e)
			}
		}
	}
	return ans
}
