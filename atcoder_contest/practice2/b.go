package main

import (
	"bufio"
	. "bytes"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func add(arg ...int) (sum int) {
	for i := range arg {
		sum += arg[i]
	}
	return
}

type RMQ struct {
	n    int
	x    []int
	unit int
	op   func(x ...int) int
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

func (rmq *RMQ) Debug() {
	for i := 0; i < rmq.n; i++ {
		if i > 0 {
			fmt.Fprintf(wr, " ")
		}
		fmt.Fprintf(wr, "%v", rmq.x[i+rmq.n])
	}
	fmt.Fprintf(wr, "\n")
}

var bytes []byte
var l int

func parseInt() (result int) {
	for bytes[l] < '0' || bytes[l] > '9' {
		l++
	}
	for '0' <= bytes[l] && bytes[l] <= '9' {
		result = result*10 + int(bytes[l]-'0')
		l++
	}
	return result
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)

	var buf Buffer
	buf.ReadFrom(os.Stdin)
	bytes = buf.Bytes()

	n := parseInt()
	q := parseInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = parseInt()
	}
	tree := RMQ{}
	tree.Create(a)

	for i := 0; i < q; i++ {
		t := parseInt()
		switch t {
		case 0:
			p, x := parseInt(), parseInt()
			tree.Update(p, tree.Query(p, p+1)+x)
		case 1:
			l, r := parseInt(), parseInt()
			wr.WriteString(strconv.Itoa(tree.Query(l, r)))
			wr.WriteByte('\n')
		}
	}

}
