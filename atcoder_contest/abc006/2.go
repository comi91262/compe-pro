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

	dp := make([]int, 1000001)

	dp[1] = 0
	dp[2] = 0
	dp[3] = 1
	for i := 4; 3 < n && i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + dp[i-3]) % 10007
	}

	fmt.Fprintf(writer, "%v\n", dp[n])
}

