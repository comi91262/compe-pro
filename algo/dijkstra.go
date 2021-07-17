package algo

import (
	"container/heap"
)

func dijkstra(start int, dist []int, g [][]edge, queue PriorityQueue) []int {
	heap.Init(&queue)
	const inf = 1 << 60

	for i := 0; i < len(dist); i++ {
		dist[i] = inf
	}
	dist[start] = 0
	heap.Push(&queue, &Item{value: start, priority: 0})

	for queue.Len() > 0 {
		v := heap.Pop(&queue).(*Item)
		for _, e := range g[v.value] {
			if dist[e.to] > dist[v.value]+e.cost {
				dist[e.to] = dist[v.value] + e.cost
				heap.Push(&queue, &Item{value: e.to, priority: dist[e.to]})
			}
		}
	}

	return dist
}
