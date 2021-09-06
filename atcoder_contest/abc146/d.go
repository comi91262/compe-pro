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

func dfs(s, pre, pc, mx int, g [][]edge) {
	c := 0
	for nex := range g[s] {
		if g[s][nex].to == pre {
			continue
		}
		if c == pc {
			c = (c + 1) % mx
		}
		g[s][nex].cost = c
		dfs(g[s][nex].to, s, c, mx, g)
		c = (c + 1) % mx

	}
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var g = make([][]edge, n)
	var a = make([]int, n-1)
	var b = make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
		a[i]--
		b[i]--
		g[a[i]] = append(g[a[i]], edge{to: b[i], cost: 0})
		g[b[i]] = append(g[b[i]], edge{to: a[i], cost: 0})
	}
	mx := 0
	for i := range g {
		mx = max(mx, len(g[i]))
	}
	dfs(0, 0, -1, mx, g)

	fmt.Fprintf(writer, "%v\n", mx)
	for i := 0; i < n-1; i++ {
		for j := range g[a[i]] {
			if g[a[i]][j].to == b[i] {
				// fmt.Fprintf(writer, "%v %v %v\n", a[i], b[i], g[a[i]][j].cost+1)
				fmt.Fprintf(writer, "%v\n", g[a[i]][j].cost+1)
				break
			}
		}
	}

}
