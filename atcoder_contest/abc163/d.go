package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const mod = 1_000_000_000 + 7

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var k int
	fmt.Fscan(reader, &k)

	a := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		a[i] = i
	}
	for i := 0; i < n; i++ {
		a[i+1] += a[i]
	}

	ans := 0
	for i := k; i <= n+1; i++ {
		mx := a[n]
		if n-i-1 >= 0 {
			mx -= a[n-i]
		}
		mn := a[i-1]
		ans += (mx - mn + 1) % mod
	}
	fmt.Fprintf(writer, "%v\n", ans%mod)
}
