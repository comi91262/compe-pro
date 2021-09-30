package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type QItem interface {
	Less(item QItem) bool
}

type priorityQueueImpl []QItem

func (pqi priorityQueueImpl) Len() int {
	return len(pqi)
}

func (pqi priorityQueueImpl) Less(i, j int) bool {
	return pqi[i].Less(pqi[j])
}

func (pqi priorityQueueImpl) Swap(i, j int) {
	pqi[i], pqi[j] = pqi[j], pqi[i]
}

func (pqi *priorityQueueImpl) Push(x interface{}) {
	item := x.(QItem)
	*pqi = append(*pqi, item)
}

func (pqi *priorityQueueImpl) Pop() interface{} {
	old := *pqi
	n := len(old)
	item := old[n-1]
	*pqi = old[0 : n-1]
	return item
}

type PriorityQueue struct {
	priorityQueueImpl
}

func NewPriorityQueue() *PriorityQueue {
	var pq PriorityQueue
	heap.Init(&pq.priorityQueueImpl)
	return &pq
}

func (pq *PriorityQueue) Push(item QItem) {
	heap.Push(&pq.priorityQueueImpl, item)
}

func (pq *PriorityQueue) Pop() QItem {
	return heap.Pop(&pq.priorityQueueImpl).(QItem)
}

func (pq *PriorityQueue) Front() QItem {
	return pq.priorityQueueImpl[0]
}

func (pq *PriorityQueue) Length() int {
	return pq.priorityQueueImpl.Len()
}

func (i node) Less(j QItem) bool {
	return i.dist < j.(node).dist
}

type node struct {
	num  int
	dist int
}

type edge struct {
	to   int
	time int
	in   int
	cost int
}

func dijkstra(start int, g [][]edge) []int {
	queue := NewPriorityQueue()
	const inf = 1 << 60

	dist := make([]int, len(g))
	for i := 0; i < len(dist); i++ {
		dist[i] = inf
	}
	dist[start] = 0
	queue.Push(node{start, dist[start]})

	for queue.Length() > 0 {
		v := queue.Pop().(node)
		if dist[v.num] < v.dist {
			continue
		}
		for _, e := range g[v.num] {
			cost := (dist[v.num]+e.in-1)/e.in*e.in + e.time
			if dist[e.to] > cost {
				dist[e.to] = cost
				queue.Push(node{e.to, dist[e.to]})
			}
		}
	}

	return dist
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)
	var x int
	fmt.Fscan(reader, &x)
	var y int
	fmt.Fscan(reader, &y)

	var g = make([][]edge, n)
	var a = make([]int, m)
	var b = make([]int, m)
	var t = make([]int, m)
	var k = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a[i], &b[i], &t[i], &k[i])
		a[i]--
		b[i]--
		g[a[i]] = append(g[a[i]], edge{b[i], t[i], k[i], 0})
		g[b[i]] = append(g[b[i]], edge{a[i], t[i], k[i], 0})
	}

	dist := dijkstra(x-1, g)

	if dist[y-1] != 1<<60 {
		fmt.Fprintf(writer, "%v\n", dist[y-1])
	} else {
		fmt.Fprintf(writer, "%v\n", -1)
	}
}
