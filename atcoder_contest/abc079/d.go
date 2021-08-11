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

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}

func dijkstra(start int, g [][]edge) []int {
	queue := PriorityQueue{}
	heap.Init(&queue)
	const inf = 1 << 60

	dist := make([]int, len(g))
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

type edge struct {
	to   int
	cost int
}

func main() {
	defer writer.Flush()

	var h int
	fmt.Fscan(reader, &h)
	var w int
	fmt.Fscan(reader, &w)

	c := make([][]int, 10)
	for i := 0; i < 10; i++ {
		c[i] = make([]int, 10)
		for j := 0; j < 10; j++ {
			fmt.Fscan(reader, &c[i][j])
		}
	}

	var g = make([][]edge, 10)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			g[i] = append(g[i], edge{j, c[i][j]})
		}
	}

	d := make([][]int, 10)
	for i := 0; i < 10; i++ {
		d[i] = dijkstra(i, g)
	}

	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
		for j := 0; j < w; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	m := map[int]int{}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			m[a[i][j]]++
		}
	}
	ans := 0
	for k, v := range m {
		if k == -1 {
			continue
		}
		if d[k][1] == 1<<60 {
			continue
		}

		ans += d[k][1] * v
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
