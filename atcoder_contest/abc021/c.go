package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const mod = 1_000_000_000 + 7

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

var cnt [300000]int

func dijkstra(start int, dist []int, g [][]edge, queue PriorityQueue) []int {
	heap.Init(&queue)
	const inf = 1 << 60

	for i := 0; i < len(dist); i++ {
		dist[i] = inf
	}

	dist[start] = 0
	cnt[start] = 1

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
				cnt[e.to] = cnt[v.value]
			} else if dist[e.to] == dist[v.value]+e.cost {
				cnt[e.to] = cnt[e.to] + cnt[v.value]%mod
			}
		}
	}

	return dist
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var m int
	fmt.Fscan(reader, &m)

	var g = make([][]edge, n+1)
	for i := 1; i < m+1; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		g[x] = append(g[x], edge{to: y, cost: 1})
		g[y] = append(g[y], edge{to: x, cost: 1})
	}
	q := make(PriorityQueue, 0)
	dist := make([]int, 300000)

	dist = dijkstra(a, dist, g, q)

	fmt.Fprintf(writer, "%v\n", cnt[b]%mod)
}
