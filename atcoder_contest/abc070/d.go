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

type edge struct {
	to   int
	cost int
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var g = make([][]edge, n+1)
	for i := 1; i < n; i++ {
		var a, b, c int
		fmt.Fscan(reader, &a, &b, &c)
		g[a] = append(g[a], edge{b, c})
		g[b] = append(g[b], edge{a, c})
	}

	var q int
	fmt.Fscan(reader, &q)
	var k int
	fmt.Fscan(reader, &k)

	d := dijkstra(k, g)

	for i := 0; i < q; i++ {
		var x, y int
		fmt.Fscan(reader, &x)
		fmt.Fscan(reader, &y)

		fmt.Fprintf(writer, "%v\n", d[x]+d[y])
	}
}
