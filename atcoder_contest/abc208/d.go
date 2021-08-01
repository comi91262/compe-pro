package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)
	var a = make([]int, m+1)
	var b = make([]int, m+1)
	var c = make([]int, m+1)
	for i := 1; i < m+1; i++ {
		fmt.Fscan(reader, &a[i])
		fmt.Fscan(reader, &b[i])
		fmt.Fscan(reader, &c[i])
	}

	var d = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		d[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			d[i][j] = 1 << 60
		}
	}
	for i := 1; i < m+1; i++ {
		d[a[i]][b[i]] = c[i]
	}

	ans := 0
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
		for s := 1; s <= n; s++ {
			for t := 1; t <= n; t++ {
				if d[s][t] == 1<<60 {
					continue
				}
				if s == t {
					continue
				}
				ans += d[s][t]
			}

		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
