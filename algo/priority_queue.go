package main

import (
	"container/heap"
)

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

type Edge struct {
	to   int
	cost int
}

func main() {
	var pq = make(PriorityQueue, 1)
	pq[0] = &Item{
		value: 0, priority: 0,
	}
	heap.Init(&pq)

	v := heap.Pop(&pq).(*Item).value
	heap.Push(&pq, &Item{value: 0, priority: 0})
}
