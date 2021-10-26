package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func scanInt() int { sc.Scan(); return parseInt(sc.Bytes()) }
func scanInts(n int) (ints []int) {
	ints = make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = scanInt()
	}
	return
}

func parseInt(b []byte) (n int) {
	for _, ch := range b {
		ch -= '0'
		if ch <= 9 {
			n = n*10 + int(ch)
		}
	}
	if b[0] == '-' {
		return -n
	}
	return
}

func add(arg ...int) (sum int) {
	for i := range arg {
		sum += arg[i]
	}
	return
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 10000), 1001001)

	n := scanInt()
	p := scanInts(n)
	m := add(p...)

	dp := make([]bool, m+1)

	dp[0] = true
	for i := 0; i < n; i++ {
		for j := m; j >= 0; j-- {
			if j+p[i] <= m {
				dp[j+p[i]] = dp[j+p[i]] || dp[j]
			}
		}
	}
	ans := 0
	for i := 0; i <= m; i++ {
		if dp[i] {
			ans++
		}
	}
	fmt.Fprintf(wr, "%v\n", ans)
}
