package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

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
		r.data[i] = max(r.data[i*2+1], r.data[i*2+2])
	}
}

func (r *RMQ) Get(i int) int {
	i += r.n - 1
	return r.data[i]
}

// [a,b) での最小の要素を取得。O(log(n))
func (r *RMQ) Query(a, b int) int { return r.querySub(a, b, 0, 0, r.n) }
func (rmq *RMQ) querySub(a, b, k, l, r int) int {
	if r <= a || b <= l {
		return 0
	} else if a <= l && r <= b {
		return rmq.data[k]
	} else {
		vl := rmq.querySub(a, b, k*2+1, l, (l+r)/2)
		vr := rmq.querySub(a, b, k*2+2, (l+r)/2, r)
		return max(vl, vr)
	}
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var t int
	fmt.Fscan(reader, &t)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	tree := RMQ{}
	tree.Create(n)
	for i := 0; i < n; i++ {
		tree.Update(i, a[i])
	}

	cnt := 0
	mx := 0
	for i := 0; i < n-1; i++ {
		l := a[i] * t / 2
		g := tree.Query(i+1, n) * t / 2
		if mx <= g-l {
			if mx == g-l {
				cnt++
			} else {
				mx = g - l
				cnt = 1
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", cnt)
}
