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

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	dp := make([]int, n+1)

	dp[0] = 0
	for i := 1; i < n; i++ {
		if i == 1 {
			dp[i] = abs(a[i] - a[i-1])
			continue
		}
		dp[i] = min(dp[i-1]+abs(a[i-1]-a[i]), dp[i-2]+abs(a[i-2]-a[i]))
	}
	fmt.Fprintf(writer, "%v\n", dp[n-1])
}
