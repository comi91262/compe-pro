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
func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}
func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

func main() {
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n)
	fmt.Fscan(reader, &k)

	var x = make([]int, n)
	var y = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i], &y[i])
	}

	dist := make([][]int, n)
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dx := abs(x[i] - x[j])
			dy := abs(y[i] - y[j])
			dist[i][j] = dx*dx + dy*dy
		}
	}

	cost := make([]int, 1<<n)
	for i := 1; i < 1<<n; i++ {
		for j := 0; j < n; j++ {
			for l := 0; l < j; l++ {
				if ((i>>j)&1) == 1 && ((i>>l)&1) == 1 {
					cost[i] = max(cost[i], dist[j][l])
				}
			}
		}
	}

	dp := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		dp[i] = make([]int, 1<<n)
		for j := 0; j < 1<<n; j++ {
			dp[i][j] = 1 << 62
		}
	}
	dp[0][0] = 0
	for i := 1; i <= k; i++ {
		for bit := 1; bit < 1<<n; bit++ {
			for b := bit; b != 0; b = (b - 1) & bit {
				dp[i][bit] = min(dp[i][bit], max(dp[i-1][bit-b], cost[b]))
			}
		}
	}

	fmt.Fprintf(writer, "%v\n", dp[k][(1<<n-1)])
}
