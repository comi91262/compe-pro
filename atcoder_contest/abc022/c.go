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

func warshallFloyd(d [][]int, n int) [][]int {
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				d[i][j] = min(d[i][j], d[i][k]+d[k][j])
			}
		}
	}

	return d
}

type edge struct {
	to   int
	cost int
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)
	var a = make([]int, m)
	var b = make([]int, m)
	var l = make([]int, m)
	var d = make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			d[i][j] = 1 << 60
		}
	}

	c := []int{}
	e := []int{}
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a[i], &b[i], &l[i])
		a[i]--
		b[i]--
		if a[i] == 0 {
			c = append(c, b[i])
			e = append(e, l[i])
			continue
		}
		d[a[i]][b[i]] = l[i]
		d[b[i]][a[i]] = l[i]
	}
	d = warshallFloyd(d, n)

	ans := 1 << 60
	for i := range c {
		for j := range c {
			if i == j {
				continue
			}
			ans = min(ans, d[c[i]][c[j]]+e[i]+e[j])

		}

	}
	if ans == 1<<60 {
		fmt.Fprintf(writer, "%v\n", -1)
	} else {
		fmt.Fprintf(writer, "%v\n", ans)
	}
}

