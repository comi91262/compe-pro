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

	var s string
	fmt.Fscan(reader, &s)
	const d = 1_000_000_000 + 7

	var dp [9][100001]int

	dp[0][0] = 1
	for i := 0; i < 9; i++ {
		for j := 0; j < len(s); j++ {
			dp[i][j+1] += dp[i][j]
			if s[j] == 'c' && i == 0 {
				dp[i+1][j+1] += dp[i][j]
			}
			if s[j] == 'h' && i == 1 {
				dp[i+1][j+1] += dp[i][j]
			}
			if s[j] == 'o' && i == 2 {
				dp[i+1][j+1] += dp[i][j]
			}
			if s[j] == 'k' && i == 3 {
				dp[i+1][j+1] += dp[i][j]
			}
			if s[j] == 'u' && i == 4 {
				dp[i+1][j+1] += dp[i][j]
			}
			if s[j] == 'd' && i == 5 {
				dp[i+1][j+1] += dp[i][j]
			}
			if s[j] == 'a' && i == 6 {
				dp[i+1][j+1] += dp[i][j]
			}
			if s[j] == 'i' && i == 7 {
				dp[i+1][j+1] += dp[i][j]
			}
			dp[i][j+1] %= d
			if i < 8 {
				dp[i+1][j+1] %= d
			}
		}

	}
	fmt.Fprintf(writer, "%v\n", dp[8][len(s)])
}
