package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type edge struct {
	to   int
	cost int
}

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}
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

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	var g = make([][]edge, n)
	var x = make([]int, m)
	var y = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &x[i], &y[i])
		x[i]--
		y[i]--
		g[x[i]] = append(g[x[i]], edge{y[i], 0})
	}
	b := make([]int, n)
	for i := 0; i < n; i++ {
		b[i] = 1 << 60
	}
	for i := 0; i < n; i++ {
		for _, j := range g[i] {
			b[j.to] = min(b[j.to], b[i], a[i])
		}
	}

	ans := -1 << 60
	for i := 0; i < n; i++ {
		ans = max(ans, a[i]-b[i])
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
