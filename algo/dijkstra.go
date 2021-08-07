package algo

import (
	"container/heap"
)

func dijkstra(start int, g [][]edge) []int {
	queue := PriorityQueue{}
	heap.Init(&queue)
	const inf = 1 << 60

	dist := make([]int, len(g))
	for i := 1; i < len(dist); i++ {
		dist[i] = inf
	}
	dist[start] = 0
	heap.Push(&queue, &Item{value: start, priority: 0})

	for queue.Len() > 0 {
		v := heap.Pop(&queue).(*Item)
		if dist[v.value] < v.priority {
			continue
		}
		for _, e := range g[v.value] {
			if dist[e.to] > dist[v.value]+e.cost {
				dist[e.to] = dist[v.value] + e.cost
				heap.Push(&queue, &Item{value: e.to, priority: dist[e.to]})
			}
		}
	}

	return dist
}
