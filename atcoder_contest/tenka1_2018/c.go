package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

type Deque struct {
	cl   int
	cr   int
	data []int
}

func (d *Deque) PushFront(x int) {
	d.cl--
	d.data[d.cl] = x
}

func (d *Deque) PushBack(x int) {
	d.data[d.cr] = x
	d.cr++
}

func (d *Deque) PopFront() int {
	v := d.data[d.cl]
	d.cl++
	return v
}

func (d *Deque) PopBack() int {
	v := d.data[d.cr-1]
	d.cr--
	return v
}

func (d *Deque) Size() int {
	return d.cr - d.cl
}

func (d *Deque) Empty() bool {
	return d.Size() == 0
}

func (d *Deque) Get(x int) int {
	return d.data[d.cl+x-1]
}

func (d *Deque) Debug() (a []int) {
	a = make([]int, d.Size())
	for i := d.cl; i < d.cr; i++ {
		a[i-d.cl] = d.data[i]
	}
	return
}

func NewDeque(n int) *Deque {
	return &Deque{
		cl:   n,
		cr:   n,
		data: make([]int, n*2, n*2),
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 10000), 1001001)
	n := scanInt()
	a := scanInts(n)
	sort.Ints(a)

	b := NewDeque(n)
	for i := 0; i < n; i++ {
		b.PushBack(a[i])
	}
	d := NewDeque(100020)
	d.PushFront(b.PopBack())

	ans := 0
	flag := true
	for !b.Empty() {
		if flag {
			if b.Size() != 1 {
				d.PushFront(b.PopFront())
				d.PushBack(b.PopFront())
			} else {
				d.PushFront(b.PopFront())
			}
		} else {
			if b.Size() != 1 {
				d.PushFront(b.PopBack())
				d.PushBack(b.PopBack())
			} else {
				d.PushFront(b.PopBack())
			}
		}
		flag = !flag
	}

	pre := d.PopBack()
	for !d.Empty() {
		nex := d.PopBack()
		ans += abs(pre - nex)
		pre = nex
	}
	for i := 0; i < n; i++ {
		b.PushFront(a[i])
	}
	d.PushFront(b.PopBack())

	flag = true
	for !b.Empty() {
		if flag {
			if b.Size() != 1 {
				d.PushFront(b.PopFront())
				d.PushBack(b.PopFront())
			} else {
				d.PushFront(b.PopFront())
			}
		} else {
			if b.Size() != 1 {
				d.PushFront(b.PopBack())
				d.PushBack(b.PopBack())
			} else {
				d.PushFront(b.PopBack())
			}
		}
		flag = !flag
	}

	sum := 0
	pre = d.PopBack()
	for !d.Empty() {
		nex := d.PopBack()
		sum += abs(pre - nex)
		pre = nex
	}
	ans = max(ans, sum)
	fmt.Fprintf(wr, "%v\n", ans)

}
