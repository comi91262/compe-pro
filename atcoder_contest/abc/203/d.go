package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func cumulativeSum2(a [][]int) [][]int {
	m := len(a) + 1

	r := make([][]int, m)
	for i := 0; i < m; i++ {
		r[i] = make([]int, m)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < m; j++ {
			r[i][j] = r[i][j-1] + a[i-1][j-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < m; j++ {
			r[i][j] += r[i-1][j]
		}
	}

	return r
}

func main() {
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	var a = make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}

	ng := -1
	ok := 1 << 60

	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2

		var b = make([][]int, n)
		for i := 0; i < n; i++ {
			b[i] = make([]int, n)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if mid < a[i][j] {
					b[i][j]++
				}
			}
		}

		c := cumulativeSum2(b)
		flag := false

		for i := 0; i < n-k+1; i++ {
			for j := 0; j < n-k+1; j++ {
				if c[i+k][j+k]-c[i+k][j]-c[i][j+k]+c[i][j] < k*k/2+1 {
					flag = true
				}
			}
		}

		if flag {
			ok = mid
		} else {
			ng = mid
		}
	}

	fmt.Fprintf(writer, "%v\n", ok)
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}
