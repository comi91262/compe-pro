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
func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func main() {
	defer writer.Flush()
	var h int
	fmt.Fscan(reader, &h)
	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		fmt.Fscan(reader, &b[i])
	}

	var dp = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, 50000)
		for j := 0; j < 50000; j++ {
			dp[i][j] = 1 << 60
		}
	}

	dp[0][0] = 0
	for i := 1; i < n+1; i++ {
		for j := 0; j < 50000; j++ {
			if j-a[i-1] >= 0 {
				dp[i][j] = min(dp[i][j-a[i-1]]+b[i-1], dp[i-1][j])
				continue
			}
			dp[i][j] = dp[i-1][j]
		}
	}

	ans := 1 << 60
	for i := h; i < 50000; i++ {
		ans = min(ans, dp[n][i])

	}
	fmt.Fprintf(writer, "%v\n", ans)

}
