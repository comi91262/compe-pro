package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

var ans []int

func dfs(s, pre int, m map[int]int, c []int, g [][]int) {
	if m[c[s]] < 1 {
		ans = append(ans, s)
	}

	for _, nex := range g[s] {
		if nex == pre {
			continue
		}

		m[c[s]]++
		dfs(nex, s, m, c, g)
		m[c[s]]--
	}
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var c = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &c[i])
	}

	var g = make([][]int, n)
	var a = make([]int, n-1)
	var b = make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
		a[i]--
		b[i]--
		g[a[i]] = append(g[a[i]], b[i])
		g[b[i]] = append(g[b[i]], a[i])
	}

	m := map[int]int{}
	dfs(0, 0, m, c, g)
	sort.Ints(ans)
	for i := range ans {
		fmt.Fprintf(writer, "%v\n", ans[i]+1)
	}
}
