package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	var s = make([]int, m)
	var c = make([]int, m)
	for i := 0; i < m; i++ {
		var ss string
		fmt.Fscan(reader, &ss)
		for j, v := range strings.Split(ss, "") {
			if v == "Y" {
				s[i] += 1 << (len(ss) - j - 1)
			}
		}
		fmt.Fscan(reader, &c[i])
	}

	var dp = make([]int, 1<<n)
	for i := range dp {
		dp[i] = 1 << 60
	}
	dp[0] = 0
	for i := 0; i < 1<<n; i++ {
		for j := 0; j < m; j++ {
			dp[i|s[j]] = min(dp[i|s[j]], dp[i]+c[j])
		}

	}
	if dp[1<<n-1] == 1<<60 {
		fmt.Fprintf(writer, "%v\n", -1)
	} else {
		fmt.Fprintf(writer, "%v\n", dp[1<<n-1])
	}
}
