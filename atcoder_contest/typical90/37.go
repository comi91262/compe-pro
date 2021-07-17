package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const inf = -1 * (1 << 60)

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
	for i := 0; i < n*4; i++ {
		r.data[i] = inf
	}
	var x int = 1
	for n > x {
		x *= 2
	}
	r.n = x
}

func (r *RMQ) Update(i, x int) {
	i += r.n - 1
	r.data[i] = x
	for i > 0 {
		i = (i - 1) / 2 // parent
		r.data[i] = max(r.data[i*2+1], r.data[i*2+2])
	}
}

// [a,b) での最小の要素を取得。O(log(n))
func (r *RMQ) Query(a, b int) int { return r.querySub(a, b, 0, 0, r.n) }
func (rmq *RMQ) querySub(a, b, k, l, r int) int {
	if r <= a || b <= l {
		return inf
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

	var w, n int
	fmt.Fscan(reader, &w, &n)

	var l = make([]int, n+1)
	var r = make([]int, n+1)
	var v = make([]int, n+1)

	for i := 1; i < n+1; i++ {
		fmt.Fscan(reader, &l[i], &r[i], &v[i])
	}

	var dp = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, w+1)
		for j := 0; j <= w; j++ {
			dp[i][j] = inf
		}
	}

	var tree = make([]RMQ, n+1)
	for i := 0; i < n+1; i++ {
		tree[i] = RMQ{}
		tree[i].Create(w + 2)
	}

	dp[0][0] = 0
	tree[0].Update(0, 0)

	for i := 1; i <= n; i++ {
		for j := 0; j <= w; j++ {
			cl := max(0, j-r[i])
			cr := max(0, j-l[i]+1)

			if cl == cr {
				dp[i][j] = dp[i-1][j]
			} else {
				val := tree[i-1].Query(cl, cr)
				if val != inf {
					dp[i][j] = max(val+v[i], dp[i-1][j])
				} else {
					dp[i][j] = dp[i-1][j]
				}
			}
		}

		for j := 0; j <= w; j++ {
			tree[i].Update(j, dp[i][j])
		}
	}

	if dp[n][w] == inf {
		fmt.Fprintf(writer, "%v\n", -1)
	} else {
		fmt.Fprintf(writer, "%v\n", dp[n][w])
	}
}
