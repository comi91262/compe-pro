package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const mod = 1_000_000_000 + 7

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	var a = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &a[i])
	}

	var dp = make([]int, n+1)
	for i := 0; i < m; i++ {
		dp[a[i]] = -1
	}

	dp[0] = 1
	for i := 0; i < n; i++ {
		if dp[i] == -1 {
			continue
		}

		if dp[i+1] != -1 {
			dp[i+1] += (dp[i] % mod)
		}
		if i+2 < n+1 && dp[i+2] != -1 {
			dp[i+2] += (dp[i] % mod)
		}
	}
	fmt.Fprintf(writer, "%v\n", (dp[n] % mod))

}
