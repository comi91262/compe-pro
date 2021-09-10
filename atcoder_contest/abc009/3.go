package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Refered to https://atcoder.jp/contests/abc009/submissions/174401

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

func isOk(ans, s []string, rest []int, n, k int) bool {
	ng := 0
	now := [26]int{}

	for i := 0; i < 26; i++ {
		now[i] = rest[i]
	}
	for i := 0; i < n; i++ {
		if ans[i] != "*" && ans[i] != s[i] {
			ng++
		}
	}
	for i := 0; i < n; i++ {
		if ans[i] == "*" {
			if now[s[i][0]-"a"[0]] > 0 {
				now[s[i][0]-"a"[0]]--
			} else {
				ng++
			}
		}
	}

	return ng <= k
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var k int
	fmt.Fscan(reader, &k)
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	ans := make([]string, n)
	for i := 0; i < n; i++ {
		ans[i] = "*"
	}

	rest := make([]int, 26)
	for i := 0; i < 26; i++ {
		rest[i] = 0
	}
	for i := 0; i < n; i++ {
		rest[s[i][0]-"a"[0]]++
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 26; j++ {
			if rest[j] > 0 {
				rest[j]--
				ans[i] = string("a"[0] + byte(j))
				if isOk(ans, s, rest, n, k) {
					break
				}
				rest[j]++
			}
		}
	}

	fmt.Fprintf(writer, "%v\n", strings.Join(ans, ""))
}
