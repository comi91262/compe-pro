package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var dp [100001]int
var g [][]int
var a, b []int

func dfs(n, pre int) {
	dp[n]++

	for _, next := range g[n] {
		if pre == next {
			continue
		}

		dfs(next, n)
		dp[n] += dp[next]
	}

}

func min(x, y int) int {
	fx := float64(x)
	fy := float64(y)
	return int(math.Min(fx, fy))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	g = make([][]int, n+1)
	a = make([]int, n+1)
	b = make([]int, n+1)

	for i := 1; i < n; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
		g[a[i]] = append(g[a[i]], b[i])
		g[b[i]] = append(g[b[i]], a[i])
	}

	dfs(1, 1)
	sum := 0
	for i := 1; i < n; i++ {
		p := min(dp[a[i]], dp[b[i]])
		sum += p * (n - p)
	}
	fmt.Fprintf(writer, "%d\n", sum)
}
