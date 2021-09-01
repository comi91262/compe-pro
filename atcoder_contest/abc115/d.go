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
	var x int
	fmt.Fscan(reader, &x)

	dp := make([]int, n+1)
	bl := make([]int, n+1)
	dp[0] = 1
	bl[0] = 1
	for i := 1; i <= n; i++ {
		dp[i] = 2*dp[i-1] + 3
		bl[i] = 2*bl[i-1] + 1
	}

	ans := 0
	for n > 0 {
		switch {
		case x <= 1:
			goto L
		case x <= 1+dp[n-1]:
			if n-1 == 0 {
				ans++
			}
			x--
			n--
			break
		case x <= 1+dp[n-1]+1:
			ans += 1 + bl[n-1]
			goto L
		case x <= 1+dp[n-1]+1+dp[n-1]:
			if n-1 == 0 {
				ans++
			}
			ans += 1 + bl[n-1]
			x -= (1 + dp[n-1] + 1)
			n--
			break
		case x <= 1+dp[n-1]+1+dp[n-1]+1:
			ans += 1 + bl[n-1]*2
			goto L
		}
	}
L:
	fmt.Fprintf(writer, "%v\n", ans)
}
