package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type Item struct {
	value    int
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

type edge struct {
	to   int
	cost int
}

var d1 [100001]int
var dn [100001]int

const INF = 1 << 60

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

func main() {
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	var g = make([][]edge, n+1)
	for i := 0; i < m; i++ {
		var a, b, c int
		fmt.Fscan(reader, &a, &b, &c)
		g[a] = append(g[a], edge{to: b, cost: c})
		g[b] = append(g[b], edge{to: a, cost: c})
	}

	q := make(PriorityQueue, 0)

	d := make([]int, 100001)
	dijkstra(1, d, g, q)
	for i := 1; i < n+1; i++ {
		d1[i] = d[i]
	}

	dijkstra(n, d, g, q)
	for i := 1; i < n+1; i++ {
		dn[i] = d[i]
	}

	for i := 1; i < n+1; i++ {
		fmt.Fprintf(writer, "%d\n", d1[i]+dn[i])
	}
}
