package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	var b = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
	}

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j < m+1; j++ {
			dp[i][j] = 1 << 60
		}
	}

	dp[0][0] = 0
	for i := 0; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			if i < n {
				dp[i+1][j] = min(dp[i+1][j], dp[i][j]+1)
			}
			if j < m {
				dp[i][j+1] = min(dp[i][j+1], dp[i][j]+1)
			}
			if i < n && j < m {
				if a[i] == b[j] {
					dp[i+1][j+1] = min(dp[i+1][j+1], dp[i][j])
				} else {
					dp[i+1][j+1] = min(dp[i+1][j+1], dp[i][j]+1)
				}
			}
		}
	}
	// printer(dp, 1<<60)

	fmt.Fprintf(writer, "%v\n", dp[n][m])
}
func printer(a [][]int, limit int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			if a[i][j] >= limit {
				fmt.Fprintf(writer, "%5v ", "x")
				continue
			}
			fmt.Fprintf(writer, "%5v ", a[i][j])
		}
		fmt.Fprintf(writer, "\n")
	}
}
