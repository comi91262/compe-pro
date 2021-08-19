package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

// Library Check 検証済
// RMQ(Range Min Query): [0, n) について、最小値を管理する構造体
type RMQ struct {
	n    int
	data []int
}

func (r *RMQ) Create(n int) {
	r.data = make([]int, n*4)
	var x int = 1
	for n > x {
		x *= 2
	}
	r.n = x
}

// i 番目の要素をxに更新。O(log(n))
func (r *RMQ) Update(i, x int) {
	i += r.n - 1
	r.data[i] = x
	for i > 0 {
		i = (i - 1) / 2 // parent
		r.data[i] = r.data[i*2+1] ^ r.data[i*2+2]
	}
}

func (r *RMQ) Get(i int) int {
	i += r.n - 1
	return r.data[i]
}

func (r *RMQ) Query(a, b int) int { return r.querySub(a, b, 0, 0, r.n) }
func (rmq *RMQ) querySub(a, b, k, l, r int) int {
	if r <= a || b <= l {
		return 0
	} else if a <= l && r <= b {
		return rmq.data[k]
	} else {
		vl := rmq.querySub(a, b, k*2+1, l, (l+r)/2)
		vr := rmq.querySub(a, b, k*2+2, (l+r)/2, r)
		return vl ^ vr
	}
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var q int
	fmt.Fscan(reader, &q)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	r := &RMQ{}
	r.Create(n)

	for i := 0; i < n; i++ {
		r.Update(i, a[i])
	}

	for i := 0; i < q; i++ {
		var t, x, y int
		fmt.Fscan(reader, &t, &x, &y)

		switch t {
		case 1:
			x--
			r.Update(x, r.Get(x)^y)
		case 2:
			x--
			y--
			fmt.Fprintf(writer, "%v\n", r.Query(x, y+1))
		}
	}
}
