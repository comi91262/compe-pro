package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var a int
	fmt.Fscan(reader, &a)

	var x = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
	}

	ans := 0
	for i := 1; i <= n; i++ {
		b := a * i

		dp := make([][][]int, n+1)
		for j := 0; j < n+1; j++ {
			dp[j] = make([][]int, b+1)
			for k := 0; k < b+1; k++ {
				dp[j][k] = make([]int, i+1)
			}
		}
		dp[0][0][0] = 1
		for j := 0; j < n; j++ {
			for k := 0; k <= b; k++ {
				for l := 0; l <= i; l++ {
					dp[j+1][k][l] += dp[j][k][l]
					if k >= x[j] && i >= l+1 {
						dp[j+1][k][l+1] += dp[j][k-x[j]][l]
					}
				}
			}
		}

		ans += dp[n][b][i]
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
