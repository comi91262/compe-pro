package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func dfs(s, e, pre, d int, g [][]int) bool {
	if d == 2 {
		return s == e
	}

	flag := false
	for _, nex := range g[s] {
		if nex == pre {
			continue
		}

		if dfs(nex, e, s, d+1, g) {
			flag = true
		}
	}

	return flag
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	var g = make([][]int, n)
	var a = make([]int, m)
	var b = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
		a[i]--
		b[i]--
		g[a[i]] = append(g[a[i]], b[i])
		g[b[i]] = append(g[b[i]], a[i])
	}

	if dfs(0, n-1, 0, 0, g) {
		fmt.Fprintf(writer, "%v\n", "POSSIBLE")
	} else {
		fmt.Fprintf(writer, "%v\n", "IMPOSSIBLE")
	}
}
