package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func pow(a, x int) int {
	r := 1
	for x > 0 {
		if x&1 == 1 {
			r *= a
		}
		a *= a
		x >>= 1
	}
	return r
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

	var a = []int{1}
	for i := 1; i < 7; i++ {
		a = append(a, pow(6, i))
	}
	for i := 1; i < 6; i++ {
		a = append(a, pow(9, i))
	}

	sort.Ints(a)

	max := 100001
	// max := n + 1

	var dp = make([][]int, 13)
	for i := 0; i < 13; i++ {
		dp[i] = make([]int, max)
		for j := 0; j < max; j++ {
			dp[i][j] = 1 << 60
		}
	}

	dp[0][0] = 0
	for i := 1; i < 13; i++ {
		for j := 0; j < max; j++ {
			if j-a[i-1] >= 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-a[i-1]]+1)
				continue
			}
			dp[i][j] = dp[i-1][j]
		}
	}

	// for i := 0; i < 13; i++ {
	// 	for j := 0; j < max; j++ {
	// 		if dp[i][j] == 1<<60 {
	// 			fmt.Fprintf(writer, "%v ", "i")
	// 		} else {
	// 			fmt.Fprintf(writer, "%v ", dp[i][j])
	// 		}
	// 	}

	// 	fmt.Fprintf(writer, "\n")
	// }
	fmt.Fprintf(writer, "%v\n", dp[12][n])
}
