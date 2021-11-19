package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func scanInt() int       { sc.Scan(); return parseInt(sc.Bytes()) }
func scanString() string { sc.Scan(); return string(sc.Bytes()) }
func scanFloat() float64 { sc.Scan(); return parseFloat(sc.Bytes()) }
func scanInts(n int) (ints []int) {
	ints = make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = scanInt()
	}
	return
}

func scanPairInts(n int) (a, b []int) {
	a = make([]int, n)
	b = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
		b[i] = scanInt()
	}
	return
}

func scanStrings(n int) (st []string) {
	st = make([]string, n)
	for i := 0; i < n; i++ {
		st[i] = scanString()
	}
	return
}
func scanFloats(n int) (fs []float64) {
	fs = make([]float64, n)
	for i := 0; i < n; i++ {
		fs[i] = scanFloat()
	}
	return fs
}

func parseInt(b []byte) (n int) {
	for _, ch := range b {
		ch -= '0'
		if ch <= 9 {
			n = n*10 + int(ch)
		}
	}
	if b[0] == '-' {
		return -n
	}
	return
}

var float64pow10 = []float64{
	1e0, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6,
	1e7, 1e8, 1e9, 1e10, 1e11, 1e12,
	1e13, 1e14, 1e15, 1e16, 1e17, 1e18,
	1e19, 1e20, 1e21, 1e22,
}

func parseFloat(b []byte) float64 {
	f, dot := 0.0, 0
	for i, ch := range b {
		if ch == '.' {
			dot = i + 1
			continue
		}
		if ch -= '0'; ch <= 9 {
			f = f*10 + float64(ch)
		}
	}
	if dot != 0 {
		f /= float64pow10[len(b)-dot]
	}
	if b[0] == '-' {
		return -f
	}
	return f
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
func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

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

type Int int

func (i Int) Less(j QItem) bool {
	return i < j.(Int)
}

type queue []int

func (q *queue) push(n int) {
	*q = append(*q, n)
}

func (q *queue) pop() int {
	v := (*q)[0]
	(*q) = (*q)[1:]
	return v
}

func (q *queue) front() int {
	return (*q)[0]
}

func (q *queue) empty() bool {
	return len(*q) == 0
}
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 10000), 1001001)

	q := scanInt()
	que := queue{}
	pq := NewPriorityQueue()

	isSorted := false
	for i := 0; i < q; i++ {
		t := scanInt()
		switch t {
		case 1:
			x := scanInt()
			que.push(x)
		case 2:
			if isSorted && !pq.Empty() {
				fmt.Fprintf(wr, "%v\n", pq.Pop())
			} else {
				fmt.Fprintf(wr, "%v\n", que.pop())
			}
		case 3:
			isSorted = true
			for !que.empty() {
				v := que.pop()
				pq.Push(Int(v))
			}
		}

	}
}
