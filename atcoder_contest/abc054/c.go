package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

var ans = 0

func dfsd(s, l int, used []bool, g [][]int) {
	if used[s] {
		return
	}

	used[s] = true
	cnt := 0
	for _, n := range g[s] {
		if used[n] {
			cnt++
			continue
		}

		dfsd(n, l+1, used, g)
	}
	if cnt-len(g[s]) == 0 {
		if l+1 == len(g) {
			ans++
		}
	}
	used[s] = false
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	var g = make([][]int, n+1)
	var a = make([]int, m+1)
	var b = make([]int, m+1)
	for i := 1; i < m+1; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
		g[a[i]] = append(g[a[i]], b[i])
		g[b[i]] = append(g[b[i]], a[i])
	}

	var used = make([]bool, n+1)
	dfsd(1, 1, used, g)

	fmt.Fprintf(writer, "%v\n", ans)
}
