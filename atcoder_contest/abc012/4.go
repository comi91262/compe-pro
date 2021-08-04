package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

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

func dijkstra(start int, g [][]edge, queue PriorityQueue) []int {
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

type edge struct {
	to   int
	cost int
}

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}
func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}
func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	var g = make([][]edge, n+1)
	for i := 1; i < m+1; i++ {
		var a, b, t int
		fmt.Fscan(reader, &a, &b, &t)
		g[a] = append(g[a], edge{to: b, cost: t})
		g[b] = append(g[b], edge{to: a, cost: t})
	}

	ans := 1 << 60
	for i := 1; i <= n; i++ {
		queue := PriorityQueue{}
		dist := dijkstra(i, g, queue)
		mx := 0
		for i := range dist {
			mx = max(mx, dist[i])
		}
		ans = min(ans, mx)
	}

	fmt.Fprintf(writer, "%v\n", ans)

}
