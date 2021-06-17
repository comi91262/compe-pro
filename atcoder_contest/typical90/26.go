package main

import (
	"bufio"
	"fmt"
	"os"
)

const MAX = 100001

var g [MAX][]int
var c [MAX]int

func dfs(n, pre int) {
	c[n] = pre

	for _, next := range g[n] {
		if c[next] >= 1 {
			continue
		}
		dfs(next, 3-pre)
	}

}

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	for i := 0; i < n; i++ {
		var x, y int
		fmt.Fscan(reader, &x, &y)
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}

	dfs(1, 1)

	d0 := make([]int, 0, n+1)
	d1 := make([]int, 0, n+1)
	for i := 1; i < n+1; i++ {
		if c[i] == 1 {
			d1 = append(d1, i)
		} else {
			d0 = append(d0, i)
		}
	}

	if len(d1) > len(d0) {
		for i := 0; i < n/2; i++ {
			fmt.Fprintf(writer, "%d\n", d1[i])
		}
	} else {
		for i := 0; i < n/2; i++ {
			fmt.Fprintf(writer, "%d\n", d0[i])
		}
	}

}
