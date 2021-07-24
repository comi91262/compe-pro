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

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

// 400 500 600
func cost(pos, remain, r, g, b int) int {
	switch {
	case remain >= g+b:
		return abs(400 - pos)

	case remain >= b:
		return abs(500 - pos)

	default:
		return abs(600 - pos)
	}

}

func main() {
	defer writer.Flush()

	var r, g, b int
	fmt.Fscan(reader, &r, &g, &b)

	n := 1000
	w := g + r + b
	inf := 1 << 60

	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, w+1)
		for j := 0; j < w+1; j++ {
			dp[i][j] = inf
		}
	}

	for i := 0; i < n; i++ {
		dp[i][w] = 0
	}
	for i := 1; i < n; i++ {
		for j := 0; j < w; j++ {
			dp[i][j] = min(dp[i-1][j+1]+cost(i, j, r, g, b), dp[i-1][j])
		}
	}

	ans := inf
	for i := 0; i < n; i++ {
		ans = min(ans, dp[i][0])
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
