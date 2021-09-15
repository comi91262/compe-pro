package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type RMQ struct {
	n    int
	x    []int
	unit int
	op   func(x ...int) int
}

func add(arg ...int) (sum int) {
	for i := range arg {
		sum += arg[i]
	}
	return
}

func (rmq *RMQ) Create(seq []int) {
	rmq.n = len(seq)
	rmq.x = make([]int, len(seq)*2)
	rmq.unit = 0
	rmq.op = add

	for i := range rmq.x {
		rmq.x[i] = rmq.unit
	}
	for i, x := range seq {
		rmq.x[i+len(seq)] = x
	}
	for i := len(seq) - 1; i > 0; i-- {
		rmq.x[i] = rmq.op(rmq.x[i<<1], rmq.x[i<<1|1])
	}
}

// i 番目の要素をxに更新。O(log(n))
func (rmq *RMQ) Update(i, x int) {
	i += rmq.n
	rmq.x[i] = x
	for i > 1 {
		i >>= 1
		rmq.x[i] = rmq.op(rmq.x[i<<1], rmq.x[i<<1|1])
	}
}

// [l,r) での最小の要素を取得。O(log(n))
func (rmq *RMQ) Query(l, r int) int {
	l += rmq.n
	r += rmq.n
	vl := rmq.unit
	vr := rmq.unit

	for l < r {
		if l&1 > 0 {
			vl = rmq.op(vl, rmq.x[l])
			l += 1
		}
		if r&1 > 0 {
			r -= 1
			vr = rmq.op(rmq.x[r], vr)
		}
		l >>= 1
		r >>= 1
	}
	return rmq.op(vl, vr)
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	tree := RMQ{}

	tree.Create(make([]int, n))

	ans := 0
	for i := 0; i < n; i++ {
		ans += tree.Query(a[i]-1, n)
		tree.Update(a[i]-1, 1)
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
