package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const mod = 998244353

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
	var n int
	fmt.Fscan(reader, &n)
	var k int
	fmt.Fscan(reader, &k)

	var l = make([]int, k)
	var r = make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &l[i])
		fmt.Fscan(reader, &r[i])
	}

	dp := make([]int, n+1)
	sum := make([]int, n+1)

	dp[0], sum[0] = 1, 1
	for i := 1; i < n; i++ {
		sum[i] += sum[i-1] + dp[i-1]
		sum[i] %= mod

		for j := 0; j < k; j++ {
			left := max(0, i-r[j])
			right := max(0, i-l[j]+1)
			dp[i] += (sum[right] - sum[left] + mod)
			dp[i] %= mod
		}
	}
	fmt.Fprintf(writer, "%v\n", dp[n-1])
}
