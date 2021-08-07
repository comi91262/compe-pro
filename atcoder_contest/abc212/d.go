package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}

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

func main() {
	defer writer.Flush()

	var q int
	fmt.Fscan(reader, &q)

	var x = make([]int, q)
	var p = make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &p[i])
		switch p[i] {
		case 1:
			fmt.Fscan(reader, &x[i])
		case 2:
			fmt.Fscan(reader, &x[i])
		case 3:
		}

	}

	r := PriorityQueue{}
	heap.Init(&r)

	t := 0
	for i := 0; i < q; i++ {
		//	for j := 0; j < r.Len(); j++ {
		//		fmt.Fprintf(writer, "%v %v\n", j, r[j])
		//	}
		switch p[i] {
		case 1:
			heap.Push(&r, &Item{value: x[i] - t, priority: x[i] - t})
		case 2:
			t += x[i]
		case 3:
			fmt.Fprintf(writer, "%v\n", heap.Pop(&r).(*Item).value+t)
		}

	}

}
