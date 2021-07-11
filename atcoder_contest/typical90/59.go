package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n, m, q int
	fmt.Fscan(reader, &n, &m, &q)

	x := make([]int, m)
	y := make([]int, m)

	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &x[i])
		fmt.Fscan(reader, &y[i])
		x[i]--
		y[i]--
	}

	g := make([]int, m)
	for i := 0; i < m; i++ {
		g[i] = y[i]*n + x[i]
	}

	sort.Ints(g)
	for i := 0; i < m; i++ {
		x[i] = g[i] % n
		y[i] = g[i] / n
	}

	var a = make([]int, q)
	var b = make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &a[i])
		fmt.Fscan(reader, &b[i])
		a[i]--
		b[i]--
	}

	for i := 0; i*64 < q; i++ {
		dp := make([]uint, n)
		for j := i * 64; j < (i+1)*64 && j < q; j++ {
			dp[a[j]] |= 1 << (j - i*64)
		}
		for j := 0; j < m; j++ {
			dp[y[j]] |= dp[x[j]]
		}
		for j := i * 64; j < (i+1)*64 && j < q; j++ {
			if dp[b[j]]>>(j-i*64)&1 == 1 {
				fmt.Fprintf(writer, "%v\n", "Yes")
			} else {
				fmt.Fprintf(writer, "%v\n", "No")
			}
		}
	}
}
