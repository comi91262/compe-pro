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
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var c int
	fmt.Fscan(reader, &c)

	dp := make([][][]float64, 101)
	for i := 0; i <= 100; i++ {
		dp[i] = make([][]float64, 101)
		for j := 0; j <= 100; j++ {
			dp[i][j] = make([]float64, 101)
		}
	}
	dp[99][99][99] = 1.0
	for i := 99; i >= 0; i-- {
		for j := 99; j >= 0; j-- {
			for k := 99; k >= 0; k-- {
				if i+j+k == 0 {
					continue
				}
				now := 0.0
				now += (dp[i+1][j][k]) * float64(i)
				now += (dp[i][j+1][k]) * float64(j)
				now += (dp[i][j][k+1]) * float64(k)
				dp[i][j][k] = now/float64(i+j+k) + 1
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", dp[a][b][c])
}
