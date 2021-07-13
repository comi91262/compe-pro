package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	var s string
	fmt.Fscan(reader, &n, &s)
	const d = 1_000_000_000 + 7

	var dp [8][100001]int

	dp[0][0] = 1
	for i := 0; i < 8; i++ {
		for j := 0; j < len(s); j++ {
			dp[i][j+1] += dp[i][j]
			if s[j] == 'a' && i == 0 {
				dp[i+1][j+1] += dp[i][j]
			}
			if s[j] == 't' && i == 1 {
				dp[i+1][j+1] += dp[i][j]
			}
			if s[j] == 'c' && i == 2 {
				dp[i+1][j+1] += dp[i][j]
			}
			if s[j] == 'o' && i == 3 {
				dp[i+1][j+1] += dp[i][j]
			}
			if s[j] == 'd' && i == 4 {
				dp[i+1][j+1] += dp[i][j]
			}
			if s[j] == 'e' && i == 5 {
				dp[i+1][j+1] += dp[i][j]
			}
			if s[j] == 'r' && i == 6 {
				dp[i+1][j+1] += dp[i][j]
			}
			dp[i][j+1] %= d
			if i < 7 {
				dp[i+1][j+1] %= d
			}
		}

	}
	fmt.Fprintf(writer, "%v\n", dp[7][len(s)])
}
