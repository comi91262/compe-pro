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

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var l int
	fmt.Fscan(reader, &l)
	var x = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
	}
	var t1 int
	fmt.Fscan(reader, &t1)
	var t2 int
	fmt.Fscan(reader, &t2)
	var t3 int
	fmt.Fscan(reader, &t3)

	dp := make([]int, l+10)
	for j := 0; j < l+10; j++ {
		dp[j] = 1 << 60
	}
	h := make([]int, l+10)
	for i := 0; i < n; i++ {
		h[x[i]] = 1
	}

	dp[0] = 0
	for j := 0; j < l+5; j++ {
		dp[j+1] = min(dp[j]+t1+t3*h[j+1], dp[j+1])
		dp[j+2] = min(dp[j]+t1+t2+t3*h[j+2], dp[j+2])
		dp[j+4] = min(dp[j]+t1+3*t2+t3*h[j+4], dp[j+4])
	}

	ans := dp[l]
	if l-1 >= 0 {
		ans = min(ans, dp[l-1]+(t1+t2)/2)
	}
	if l-2 >= 0 {
		ans = min(ans, dp[l-2]+(t1+t2)/2+t2)
	}
	if l-3 >= 0 {
		ans = min(ans, dp[l-3]+(t1+t2)/2+t2*2)
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
