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
	return pq[i].priority > pq[j].priority
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

	var k int
	fmt.Fscan(reader, &k)
	var t int
	fmt.Fscan(reader, &t)
	var a = make([]int, t)
	for i := 0; i < t; i++ {
		fmt.Fscan(reader, &a[i])
	}
	pq := PriorityQueue{}
	heap.Init(&pq)
	for i := 0; i < t; i++ {
		heap.Push(&pq, &Item{value: a[i], priority: a[i]})
	}
	for pq.Len() > 1 {
		s := heap.Pop(&pq).(*Item).value
		t := heap.Pop(&pq).(*Item).value
		if s-t > 0 {
			heap.Push(&pq, &Item{value: s - t, priority: s - t})
		}
	}
	if pq.Len() > 0 {
		s := heap.Pop(&pq).(*Item).value
		fmt.Fprintf(writer, "%v\n", s-1)
	} else {
		fmt.Fprintf(writer, "%v\n", 0)
	}
}
