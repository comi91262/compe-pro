package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type node struct {
	to   []int
	cost int
}

func dfs(s, pre int, g []node) {

	for _, n := range g[s].to {
		if n == pre {
			continue
		}
		g[n].cost += g[s].cost
		dfs(n, s, g)
	}
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var q int
	fmt.Fscan(reader, &q)

	var g = make([]node, n)
	var a = make([]int, n-1)
	var b = make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
		a[i]--
		b[i]--
		g[a[i]].to = append(g[a[i]].to, b[i])
		g[b[i]].to = append(g[b[i]].to, a[i])
		g[a[i]].cost = 0
		g[b[i]].cost = 0
	}

	var p = make([]int, q)
	var x = make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &p[i])
		fmt.Fscan(reader, &x[i])
		p[i]--
	}

	for i := 0; i < q; i++ {
		g[p[i]].cost += x[i]
	}
	dfs(0, 0, g)
	for i := 0; i < len(g); i++ {
		fmt.Fprintf(writer, "%v", g[i].cost)
		if i != len(g)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")
}
