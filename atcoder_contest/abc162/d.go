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
	var s string
	fmt.Fscan(reader, &s)

	total := 0
	t := []string{"RGB", "RBG", "GRB", "GBR", "BRG", "BGR"}
	for i := 0; i < len(t); i++ {
		total += enumatateChokudai(s, t[i])
	}

	rest := 0

	for k := 3; k <= len(s); k = k + 2 {
		for i := 0; i < len(s)-k+1; i++ {
			ss := string(s[i]) + string(s[k/2+i]) + string(s[i+k-1])
			for j := 0; j < len(t); j++ {
				if ss == t[j] {
					rest++
				}
			}
		}

	}
	fmt.Fprintf(writer, "%v\n", total-rest)
}

// 文字列sの中から, 文字列tと並ぶものを数える
// e.g  (abcbc, bc) -> 2
// O(|s||t||t|)
func enumatateChokudai(s, t string) int {
	dp := make([][]int, len(t)+1)
	for i := 0; i < len(t)+1; i++ {
		dp[i] = make([]int, len(s)+1)
	}

	dp[0][0] = 1
	for i := 0; i <= len(t); i++ {
		for j := 0; j < len(s); j++ {
			dp[i][j+1] += dp[i][j]
			for k := 0; k < len(t); k++ {
				if s[j] == t[k] && i == k {
					dp[i+1][j+1] += dp[i][j]
				}
			}
		}
	}
	return dp[len(t)][len(s)]
}
