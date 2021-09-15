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

const mod = 1_000_000_000 + 7

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

	factMod := make([]int, n)
	factMod[0] = 0
	factMod[1] = 1
	for i := 2; i < n; i++ {
		factMod[i] = i * factMod[i-1] % mod
	}

	ans := 0
	for i := 0; i < n; i++ {
		used := tree.Query(0, a[i]) + 1
		tree.Update(a[i]-1, 1)
		ans += (a[i] - used) * factMod[n-i-1]
		ans %= mod
	}
	ans++ // 辞書順x番目は 1-indexed

	fmt.Fprintf(writer, "%v\n", ans)
}
