package algo

type node struct {
	num  int
	dist int
}

func (i node) Less(j QItem) bool {
	return i.dist < j.(node).dist
}

func dijkstra(start int, g [][]edge) []int {
	const inf = 1 << 60
	queue := NewPriorityQueue()

	dist := make([]int, len(g))
	for i := 0; i < len(dist); i++ {
		dist[i] = inf
	}
	dist[start] = 0
	queue.Push(node{start, dist[start]})

	for !queue.Empty() {
		v := queue.Pop().(node)
		if dist[v.num] < v.dist {
			continue
		}
		for _, e := range g[v.num] {
			if dist[e.to] > dist[v.num]+e.cost {
				dist[e.to] = dist[v.num] + e.cost
				queue.Push(node{e.to, dist[e.to]})
			}
		}
	}
	return dist
}
