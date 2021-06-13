package main

import (
	"bufio"
	"fmt"
	"os"
)

var dp [101][100001]bool

// var dp [10][12]bool

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var t = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &t[i])
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += t[i]
	}

	dp[0][0] = true
	for i := 0; i < n; i++ {
		for j := 0; j < sum+1; j++ {
			if j >= t[i] {
				dp[i+1][j] = dp[i][j-t[i]] || dp[i][j]
			} else {
				dp[i+1][j] = dp[i][j]
			}
		}
	}

	//	fmt.Fprintf(writer, "%v\n", dp[0])
	//	fmt.Fprintf(writer, "%v\n", dp[1])
	//	fmt.Fprintf(writer, "%v\n", dp[2])
	//	fmt.Fprintf(writer, "%v\n", dp[3])

	for i := (sum + 1) / 2; i < sum+1; i++ {
		if dp[n][i] {
			fmt.Fprintf(writer, "%d\n", i)
			break
		}
	}
}

