package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod = 1_000_000_000 + 7

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)
var n, k int

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
	return c[tx][ty] - c[tx][sy-1] - c[sx-1][ty] + c[sx-1][sy-1]
}

func check(a [][]int, boader, ans int) bool {
	c := cumulativeSum2{}
	c.create(n, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if ans <= a[i][j] {
				c.add(i, j, 1)
			}
		}
	}

	c.build()

	for i := 1; i <= n-k+1; i++ {
		for j := 1; j <= n-k+1; j++ {
			if c.get(i, j, i+k-1, j+k-1) < boader {
				return false
			}
		}
	}

	return true
}

func binarySearch(a [][]int, boader int, criteria func(a [][]int, boader, answer int) bool) int {
	abs := func(x int) int {
		if x >= 0 {
			return x
		} else {
			return x * -1
		}
	}

	ng := 1 << 60
	ok := -1

	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2

		if criteria(a, boader, mid) {
			ok = mid
		} else {
			ng = mid
		}
	}

	return ok
}

func main() {
	defer writer.Flush()

	fmt.Fscan(reader, &n, &k)

	var a = make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}

	ans := binarySearch(a, k*k/2+1, check)

	fmt.Fprintf(writer, "%v\n", ans)
}
