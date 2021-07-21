package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func contains(x int, a []int) bool {
	for _, e := range a {
		if x == e {
			return true
		}
	}
	return false
}

func main() {
	defer writer.Flush()

	var n, m int
	fmt.Fscan(reader, &n, &m)

	var g = make([][]int, n+1)
	var a = make([]int, m+1)
	var b = make([]int, m+1)

	for i := 1; i <= m; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
		g[a[i]] = append(g[a[i]], b[i])
		g[b[i]] = append(g[b[i]], a[i])
	}

	ans := 0
	for i := 1; i <= n; i++ {
		group := []int{}
		group = append(group, i)
		for _, j := range g[i] {
			flag := true
			for _, k := range group {
				if !contains(k, g[j]) {
					flag = false
					break
				}

			}
			if flag {
				group = append(group, j)
			}
		}

		ans = max(ans, len(group))
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
