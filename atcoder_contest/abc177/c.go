package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const mod = 1_000_000_000 + 7

func sum(a []int) int {
	r := 0
	for i := range a {
		r += a[i]
	}
	return r
}
func pow(a, x int) int {
	r := 1
	for x > 0 {
		if x&1 == 1 {
			r = r * a % mod
		}
		a = a * a % mod
		x >>= 1
	}
	return r
}

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	s := sum(a)
	s %= mod

	s *= s
	s %= mod

	t := 0
	for i := 0; i < n; i++ {
		t += a[i] * a[i]
		t %= mod
	}

	ans := (s - t + mod) % mod
	ans *= pow(2, mod-2)

	fmt.Fprintf(writer, "%v\n", ans%mod)
}

