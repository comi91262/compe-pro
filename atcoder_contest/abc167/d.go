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
	var k int
	fmt.Fscan(reader, &k)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		a[i]--
	}

	size := 60
	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[0][i] = a[i]
	}

	for i := 1; i < size; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = dp[i-1][dp[i-1][j]]
		}
	}

	s := 0
	for i := 0; i < size; i++ {
		if k&(1<<i) != 0 {
			s = dp[i][s]
		}
	}
	fmt.Fprintf(writer, "%v\n", s+1)

}
