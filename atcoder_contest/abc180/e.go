package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

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
func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func dist(x, y, z []int, from, to int) int {
	return abs(x[to]-x[from]) + abs(y[to]-y[from]) + max(0, z[to]-z[from])
}
func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var x = make([]int, n)
	var y = make([]int, n)
	var z = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
		fmt.Fscan(reader, &y[i])
		fmt.Fscan(reader, &z[i])
	}

	dp := make([][]int, 1<<n)
	for i := 0; i < 1<<n; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = 1 << 60
		}
	}

	dp[0][0] = 0
	for bit := 0; bit < 1<<n; bit++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j {
					continue
				}
				dp[bit][i] = min(dp[bit][i], dp[bit&^(1<<i)][j]+dist(x, y, z, j, i))
			}
		}
	}
	//printer(dp, 1<<60)
	fmt.Fprintf(writer, "%v\n", dp[1<<n-1][0])
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
