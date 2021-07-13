package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const inf = 1_000_000_000 + 7

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
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

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	dp := make([][]int, 2*n)
	for i := 0; i < 2*n; i++ {
		dp[i] = make([]int, 2*n)
	}

	for i := 0; i < 2*n; i++ {
		for j := 0; j < 2*n; j++ {
			dp[i][j] = inf
		}
	}
	for i := 0; i < 2*n-1; i++ {
		dp[i][i+1] = abs(a[i] - a[i+1])
	}

	for interval := 3; interval <= 2*n; interval += 2 {
		for j := 0; j < 2*n-interval; j++ {
			cl, cr := j, j+interval

			dp[cl][cr] = min(dp[cl][cr], dp[cl+1][cr-1]+abs(a[cl]-a[cr]))

			for k := cl; k < cr; k++ {
				dp[cl][cr] = min(dp[cl][cr], dp[cl][k]+dp[k+1][cr])
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", dp[0][2*n-1])
}
