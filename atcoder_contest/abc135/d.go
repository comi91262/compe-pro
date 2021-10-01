package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const mod = 1_000_000_000 + 7

func main() {
	defer writer.Flush()

	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	a := make([]int, len(s))
	for i := 0; i < len(a); i++ {
		a[i] = int(s[i][0] - "0"[0])
	}
	var dp = make([][]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]int, 13)
	}

	dp[0][0] = 1
	for i := 0; i < len(s); i++ {
		for j := 0; j < 13; j++ {
			if s[i] == "?" {
				for k := 0; k < 10; k++ {
					rest := (j*10 + k) % 13
					dp[i+1][rest] += dp[i][j]
					dp[i+1][rest] %= mod
				}
			} else {
				rest := (j*10 + int(s[i][0]-"0"[0])) % 13
				dp[i+1][rest] += dp[i][j]
				dp[i+1][rest] %= mod
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", dp[len(s)][5])
}

