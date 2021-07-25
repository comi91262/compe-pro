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

type cumulativeSum2 [][]int

func (c *cumulativeSum2) create(n, m int) {
	*c = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		(*c)[i] = make([]int, m+1)
	}
}

func (c cumulativeSum2) add(x, y, v int) {
	c[x+1][y+1] += v
}

type point struct {
	x int
	y int
}

func pairs(n, limit int) []point {
	pair := []point{}
	for i := 1; i <= n; i++ {
		if i > limit {
			continue
		}

		y := n / i
		if i*y == n && y <= limit {
			pair = append(pair, point{i, y})
		}
	}

	return pair
}

func (c cumulativeSum2) build() {
	n := len(c)
	m := len(c[0])

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			c[i][j] += c[i][j-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			c[j][i] += c[j-1][i]
		}
	}
}

func (c cumulativeSum2) get(sx, sy, tx, ty int) int {
	return c[tx][ty] - c[tx][sy] - c[sx][ty] + c[sx][sy]
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &d[i][j])
		}
	}
	var q int
	fmt.Fscan(reader, &q)

	var p = make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &p[i])
	}

	c := cumulativeSum2{}
	c.create(n, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			c.add(i, j, d[i][j])
		}
	}
	c.build()

	mx := 0
	for i := 0; i < q; i++ {
		mx = max(mx, p[i])
	}

	r := make([][]point, mx)
	for j := 1; j <= mx; j++ {
		r[j-1] = pairs(j, n)
	}

	t := make([]int, mx)
	for i, pairs := range r {
		tmp := 0
		for _, po := range pairs {
			for j := 0; j < n-po.x+1; j++ {
				for k := 0; k < n-po.y+1; k++ {
					tmp = max(tmp, c.get(j, k, j+po.x, k+po.y))
				}
			}
		}
		t[i] = tmp
	}

	for i := 1; i < len(t); i++ {
		t[i] = max(t[i], t[i-1])
	}

	for i := 0; i < q; i++ {
		fmt.Fprintf(writer, "%v\n", t[p[i]-1])
	}
}
