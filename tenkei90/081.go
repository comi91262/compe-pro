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

	var n, k int
	fmt.Fscan(reader, &n, &k)

	var a = make([]int, n)
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
	}

	m := 5001
	c := cumulativeSum2{}
	c.create(m, m)

	for i := 0; i < n; i++ {
		c.add(a[i], b[i], 1)
	}

	c.build()

	mx := 0
	for a := 1; a <= m-k-1; a++ {
		for b := 1; b <= m-k-1; b++ {
			mx = max(mx, c.get(a-1, b-1, a+k, b+k))
		}
	}

	fmt.Fprintf(writer, "%v\n", mx)
}
