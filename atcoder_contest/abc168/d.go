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

func dijkstra(start int, g [][]edge) (dist, pre []int) {
	queue := PriorityQueue{}
	heap.Init(&queue)
	const inf = 1 << 60

	dist = make([]int, len(g))
	for i := 1; i < len(dist); i++ {
		dist[i] = inf
	}
	dist[start] = 0
	pre = make([]int, len(g))
	heap.Push(&queue, &Item{value: start, priority: 0})

	for queue.Len() > 0 {
		v := heap.Pop(&queue).(*Item)
		if dist[v.value] < v.priority {
			continue
		}
		for _, e := range g[v.value] {
			if dist[e.to] > dist[v.value]+e.cost {
				dist[e.to] = dist[v.value] + e.cost
				pre[e.to] = v.value
				heap.Push(&queue, &Item{value: e.to, priority: dist[e.to]})
			}
		}
	}

	return
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	var g = make([][]edge, n)
	var a = make([]int, m)
	var b = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
		a[i]--
		b[i]--
		g[a[i]] = append(g[a[i]], edge{to: b[i], cost: 1})
		g[b[i]] = append(g[b[i]], edge{to: a[i], cost: 1})
	}

	d, p := dijkstra(0, g)
	for i := range d {
		if d[i] == 1<<60 {
			fmt.Fprintf(writer, "%v\n", "No")
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", "Yes")
	for i := 1; i < len(p); i++ {
		fmt.Fprintf(writer, "%v\n", p[i]+1)

	}
}
