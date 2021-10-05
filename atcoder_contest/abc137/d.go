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

func (pq *PriorityQueue) Empty() bool {
	return pq.priorityQueueImpl.Len() == 0
}

func NewPriorityQueue() *PriorityQueue {
	var pq PriorityQueue
	heap.Init(&pq.priorityQueueImpl)
	return &pq
}

//func (i node) Less(j QItem) bool {
//	return i.dist < j.(node).dist
//}
type Int int

func (i Int) Less(j QItem) bool {
	return i > j.(Int)
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	var a = make([]int, n)
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
	}
	day := make([][]int, m)
	for i := 0; i < n; i++ {
		if m-a[i] >= 0 {
			day[m-a[i]] = append(day[m-a[i]], b[i])
		}
	}

	pq := NewPriorityQueue()
	ans := 0
	for i := 0; i < m; i++ {
		for j := range day[m-i-1] {
			pq.Push(Int(day[m-i-1][j]))
		}
		if !pq.Empty() {
			ans += int(pq.Pop().(Int))
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
