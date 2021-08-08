package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const mod = 998244353

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)
	var k int
	fmt.Fscan(reader, &k)

	var g = make([][]int, n+1)
	var u = make([]int, m+1)
	var v = make([]int, m+1)
	for i := 1; i < m+1; i++ {
		fmt.Fscan(reader, &u[i], &v[i])
		g[u[i]] = append(g[u[i]], v[i])
		g[v[i]] = append(g[v[i]], u[i])
	}

	var dp = make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		dp[i] = make([]int, n+1)
	}

	dp[0][1] = 1
	for i := 1; i <= k; i++ {
		total := 0
		for j := 1; j <= n; j++ {
			total += dp[i-1][j]
		}

		for j := 1; j <= n; j++ {
			dp[i][j] = total
			for _, l := range g[j] {
				dp[i][j] -= dp[i-1][l]
			}
			dp[i][j] -= dp[i-1][j]
			dp[i][j] %= mod
		}
	}

	fmt.Fprintf(writer, "%v\n", dp[k][1]%mod)
}

