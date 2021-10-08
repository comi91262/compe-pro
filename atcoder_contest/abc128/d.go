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
	return i < j.(Int)
}

type Deque struct {
	cl   int
	cr   int
	data []int
}

func (d *Deque) pushFront(x int) {
	d.cl--
	d.data[d.cl] = x
}

func (d *Deque) pushBack(x int) {
	d.data[d.cr] = x
	d.cr++
}

func (d *Deque) popFront() int {
	v := d.data[d.cl]
	d.cl++
	return v
}

func (d *Deque) popBack() int {
	v := d.data[d.cr-1]
	d.cr--
	return v
}

func (d *Deque) size() int {
	return d.cr - d.cl
}

func (d *Deque) empty() bool {
	return d.size() == 0
}

func (d *Deque) get(x int) int {
	return d.data[d.cl+x-1]
}

func create(size int) Deque {
	return Deque{cl: size, cr: size, data: make([]int, size*2, size*2)}
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
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	n := scanInt()
	k := scanInt()
	v := scanInts(n)

	ans := -1 << 60
	for i := 0; i <= k; i++ {
		for j := 0; j <= k-i; j++ {
			if i+j > min(k, n) {
				continue
			}
			for l := 0; l <= k-i-j; l++ {

				pq := NewPriorityQueue()
				d := create(n + 1)
				for i := 0; i < n; i++ {
					d.pushBack(v[i])
				}
				for a := 0; a < i; a++ {
					if !d.empty() {
						pq.Push(Int(d.popFront()))
					}
				}
				for b := 0; b < j; b++ {
					if !d.empty() {
						pq.Push(Int(d.popBack()))
					}
				}
				for c := 0; c < l; c++ {
					if !pq.Empty() {
						pq.Pop()
					}
				}
				sum := 0
				for !pq.Empty() {
					sum += int(pq.Pop().(Int))
				}
				ans = max(ans, sum)
			}
		}
	}
	fmt.Fprintf(wr, "%v\n", ans)
}
