package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var q int
	fmt.Fscan(reader, &q)
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	dp := make([]int, n)

	pre := s[0]
	for i := 1; i < n; i++ {
		dp[i] = dp[i-1]
		if pre == "A" && s[i] == "C" {
			dp[i]++
		}
		pre = s[i]
	}

	for i := 0; i < q; i++ {
		ans := 0
		var l, r int
		fmt.Fscan(reader, &l)
		fmt.Fscan(reader, &r)
		ans += dp[r-1]
		if l-1 >= 0 {
			ans -= dp[l-1]
		}
		fmt.Fprintf(writer, "%v\n", ans)
	}

}
