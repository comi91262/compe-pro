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

	var x int
	fmt.Fscan(reader, &x)

	a := []int{100, 101, 102, 103, 104, 105}
	var dp = make([][]bool, len(a)+1)
	for i := 0; i < len(a)+1; i++ {
		dp[i] = make([]bool, x+1)
	}

	dp[0][0] = true
	for i := 1; i <= len(a); i++ {
		for j := 0; j <= x; j++ {
			if j-a[i-1] >= 0 {
				dp[i][j] = dp[i][j-a[i-1]] || dp[i-1][j]
				continue
			}
			dp[i][j] = dp[i-1][j]
		}
	}
	if dp[len(a)][x] {
		fmt.Fprintf(writer, "%v\n", 1)
	} else {
		fmt.Fprintf(writer, "%v\n", 0)
	}
}
