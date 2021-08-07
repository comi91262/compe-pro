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

func reverseMatrix(a [][]int) [][]int {
	n := len(a)
	m := len(a[0])

	for i := 0; i < n/2; i++ {
		for j := 0; j < m; j++ {
			a[i][j], a[n-1-i][j] = a[n-1-i][j], a[i][j]
		}
	}

	return a
}

func main() {
	defer writer.Flush()

	var h, w, c int
	fmt.Fscan(reader, &h, &w, &c)

	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
		for j := 0; j < w; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}

	ans := 1 << 60
	for k := 0; k < 2; k++ {
		dp := make([][]int, h)
		for i := 0; i < h; i++ {
			dp[i] = make([]int, w)
			for j := 0; j < w; j++ {
				dp[i][j] = 1 << 60
			}
		}

		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				if i > 0 {
					dp[i][j] = min(dp[i-1][j], dp[i][j])
				}
				if j > 0 {
					dp[i][j] = min(dp[i][j-1], dp[i][j])
				}
				ans = min(ans, dp[i][j]+a[i][j]+c*(i+j))
				dp[i][j] = min(dp[i][j], a[i][j]-c*(i+j))
			}
		}
		a = reverseMatrix(a)
	}

	fmt.Fprintf(writer, "%v\n", ans)
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
