package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

var mod = 1000000007
var dp [1 << 18][3]int

func dfs(pos, pre int, c []string, g [][]int) {
	val1 := 1
	val2 := 1
	for _, i := range g[pos] {
		if i == pre {
			continue
		}
		dfs(i, pos, c, g)
		if c[pos] == "a" {
			val1 *= (dp[i][0] + dp[i][2])
			val2 *= (dp[i][0] + dp[i][1] + 2*dp[i][2])
		}
		if c[pos] == "b" {
			val1 *= (dp[i][1] + dp[i][2])
			val2 *= (dp[i][0] + dp[i][1] + 2*dp[i][2])
		}
		val1 %= mod
		val2 %= mod
	}

	if c[pos] == "a" {
		dp[pos][0] = val1
		dp[pos][2] = (val2 - val1 + mod) % mod
	}
	if c[pos] == "b" {
		dp[pos][1] = val1
		dp[pos][2] = (val2 - val1 + mod) % mod
	}
}
func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var c = make([]string, n+1)
	for i := 1; i < n+1; i++ {
		fmt.Fscan(reader, &c[i])
	}

	var g = make([][]int, n+1)
	var a = make([]int, n+1)
	var b = make([]int, n+1)
	for i := 1; i < n; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
		g[a[i]] = append(g[a[i]], b[i])
		g[b[i]] = append(g[b[i]], a[i])
	}

	dfs(1, -1, c, g)
	// fmt.Fprintf(writer, "%v\n", g)
	// for i := 1; i < n+1; i++ {
	// 	fmt.Fprintf(writer, "%v", i)
	// 	for j := 0; j < 3; j++ {
	// 		switch j {
	// 		case 0:
	// 			fmt.Fprintf(writer, " a: %v", dp[i][j])
	// 		case 1:
	// 			fmt.Fprintf(writer, " b: %v", dp[i][j])
	// 		case 2:
	// 			fmt.Fprintf(writer, " ab: %v", dp[i][j])
	// 		}
	// 	}
	// 	fmt.Fprintf(writer, "\n")
	// }

	fmt.Fprintf(writer, "%v", dp[1][2])
}
