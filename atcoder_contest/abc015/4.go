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

func main() {
	defer writer.Flush()
	var w int
	fmt.Fscan(reader, &w)
	var n int
	fmt.Fscan(reader, &n)
	var k int
	fmt.Fscan(reader, &k)

	var a = make([]int, n)
	var b = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i], &b[i])
	}
	var dp = make([][][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j < k+1; j++ {
			dp[i][j] = make([]int, w+1)
		}
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < k+1; j++ {
			for l := 0; l < w+1; l++ {
				if l-a[i-1] >= 0 {
					dp[i][j][l] = max(dp[i-1][j-1][l-a[i-1]]+b[i-1], dp[i-1][j][l])
				} else {
					dp[i][j][l] = dp[i-1][j][l]
				}
			}
		}
		//	for i := range dp {
		//		fmt.Fprintf(writer, "%v\n", dp[i])
		//	}
		//	fmt.Fprintf(writer, "\n")

	}

	fmt.Fprintf(writer, "%v\n", dp[n][k][w])
}
