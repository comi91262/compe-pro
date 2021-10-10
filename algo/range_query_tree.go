package algo

import "fmt"

// [0, n) について、最小値を管理する構造体
type RMQ struct {
	n    int
	x    []int
	unit int
	op   func(x ...int) int
}

func (rmq *RMQ) Create(seq []int) {
	rmq.n = len(seq)
	rmq.x = make([]int, len(seq)*2)
	rmq.unit = 1 << 60
	rmq.op = min

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

func (rmq *RMQ) Debug() {
	for i := 0; i < rmq.n; i++ {
		if i > 0 {
			fmt.Fprintf(wr, " ")
		}
		fmt.Fprintf(wr, "%v", rmq.x[i+rmq.n])
	}
	fmt.Fprintf(wr, "\n")
}
