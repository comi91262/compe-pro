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

	var d int
	fmt.Fscan(reader, &d)
	var g int
	fmt.Fscan(reader, &g)
	var p = make([]int, d)
	var c = make([]int, d)
	for i := 0; i < d; i++ {
		fmt.Fscan(reader, &p[i])
		fmt.Fscan(reader, &c[i])
		c[i] /= 100
	}

	g /= 100
	size := 400000
	dp := make([][]int, d+1)
	for i := 0; i <= d; i++ {
		dp[i] = make([]int, size)
		for j := 0; j < size; j++ {
			dp[i][j] = 1 << 60
		}
	}

	dp[0][0] = 0
	for i := 0; i < d; i++ {
		for j := 0; j < size/2; j++ {
			for k := 0; k <= p[i]; k++ {
				score := j + k*(i+1)
				if k == p[i] {
					dp[i+1][score+c[i]] = min(dp[i+1][score+c[i]], dp[i][j]+k)
					continue
				}
				dp[i+1][score] = min(dp[i+1][score], dp[i][j]+k)
			}
		}
	}
	ans := 1 << 60
	for i := g; i < size; i++ {
		ans = min(dp[d][i], ans)
	}
	//printer(dp, 1<<60)
	fmt.Fprintf(writer, "%v\n", ans)
}
func printer(a [][]int, limit int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			if a[i][j] >= limit {
				fmt.Fprintf(writer, "%2v ", "x")
				continue
			}
			fmt.Fprintf(writer, "%2v ", a[i][j])
		}
		fmt.Fprintf(writer, "\n")
	}
}
