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

func (q *queue) front() int {
	return (*q)[0]
}

func (q *queue) empty() bool {
	return len(*q) == 0
}

// ただ BFSするだけ
func bfs(start int, g [][]int) {
	q := queue{}
	q.push(start)

	for !q.empty() {
		node := q.pop()
		for _, e := range g[node] {
			q.push(e)
		}
	}
}

// 単一コストの最短経路問題を BFS でとく
func bfsMinDist(start int, g [][]int) []int {
	const inf = -(1 << 60)
	q := queue{}
	q.push(start)

	dist := make([]int, len(g)+1)

	for i := 0; i <= len(g); i++ {
		dist[i] = inf
	}
	dist[start] = 0
	for !q.empty() {
		node := q.pop()
		for _, e := range g[node] {
			if dist[e] == inf {
				dist[e] = dist[node] + 1
				q.push(e)
			}
		}
	}

	return dist
}
