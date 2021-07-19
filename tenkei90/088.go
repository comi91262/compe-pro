package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)
var ans [][][]int

const sum = 20000

func main() {
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	var g = make([][]int, n)
	x := make([]int, q)
	y := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &x[i], &y[i])
		g[x[i]-1] = append(g[x[i]-1], y[i]-1)
	}

	ans = make([][][]int, sum)
	for i := 0; i < sum; i++ {
		ans[i] = make([][]int, 2)
	}

	used := make([]int, sum)
	vec := []int{}
	dfs(0, 0, a, vec, used, g)

	for i := 0; i < sum; i++ {
		if ans[i][0] == nil || ans[i][1] == nil {
			continue
		}

		fmt.Fprintf(writer, "%v\n", len(ans[i][0]))
		for j := 0; j < len(ans[i][0]); j++ {
			if j != 0 {
				fmt.Fprintf(writer, " ")
			}
			fmt.Fprintf(writer, "%v", ans[i][0][j]+1)
		}
		fmt.Fprintf(writer, "\n")

		fmt.Fprintf(writer, "%v\n", len(ans[i][1]))
		for j := 0; j < len(ans[i][1]); j++ {
			if j != 0 {
				fmt.Fprintf(writer, " ")
			}
			fmt.Fprintf(writer, "%v", ans[i][1][j]+1)
		}
		fmt.Fprintf(writer, "\n")
	}
}

func dfs(pos, total int, a, v, used []int, g [][]int) bool {
	if pos == len(a) {
		v2 := make([]int, len(v))
		copy(v2, v)
		switch {
		case len(ans[total][0]) == 0:
			ans[total][0] = v2
		case len(ans[total][1]) == 0:
			ans[total][1] = v2
		}

		return ans[total][0] != nil && ans[total][1] != nil
	}

	if end := dfs(pos+1, total, a, v, used, g); end {
		return true
	}

	if used[pos] == 0 {
		v = append(v, pos)
		for _, e := range g[pos] {
			used[e]++
		}
		if end := dfs(pos+1, total+a[pos], a, v, used, g); end {
			return true
		}
		for _, e := range g[pos] {
			used[e]--
		}
		v = v[:len(v)-1]
	}

	return false
}
