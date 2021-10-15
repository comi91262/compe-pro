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

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	pq := PriorityQueue{}
	heap.Init(&pq)
	cnt := 0
	for i := 0; i < n; i++ {
		b := a[i]
		for b%2 == 0 {
			cnt++
			b /= 2
		}
		heap.Push(&pq, &Item{value: b, priority: b})
	}
	for i := 0; i < cnt; i++ {
		v := heap.Pop(&pq).(*Item)
		heap.Push(&pq, &Item{value: v.value * 3, priority: v.value * 3})
	}

	ans := 1 << 60
	for pq.Len() != 0 {
		v := heap.Pop(&pq).(*Item)
		ans = min(ans, v.value)
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
