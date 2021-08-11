package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func contains(x int, a []int) bool {
	for _, e := range a {
		if x == e {
			return true
		}
	}
	return false
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var ng1 int
	var ng2 int
	var ng3 int
	fmt.Fscan(reader, &ng1)
	fmt.Fscan(reader, &ng2)
	fmt.Fscan(reader, &ng3)

	ng := []int{ng1, ng2, ng3}

	var dp = make([][]bool, 102)
	for i := 0; i < 102; i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true
	for i := 0; i < 101; i++ {
		for j := 0; j < n+1; j++ {
			if j+1 < n+1 && !contains(j+1, ng) {
				dp[i+1][j+1] = dp[i+1][j+1] || dp[i][j]
			}
			if j+2 < n+1 && !contains(j+2, ng) {
				dp[i+1][j+2] = dp[i+1][j+2] || dp[i][j]
			}
			if j+3 < n+1 && !contains(j+3, ng) {
				dp[i+1][j+3] = dp[i+1][j+3] || dp[i][j]
			}
		}
	}

	for i := 0; i < 101; i++ {
		if dp[i][n] {
			fmt.Fprintf(writer, "%v\n", "YES")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "NO")

}
