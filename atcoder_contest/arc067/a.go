package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func primeFactor(n int) map[int]int {
	var m = map[int]int{}

	for i := 2; i*i <= n; i++ {
		if n%i != 0 {
			continue
		}
		for n%i == 0 {
			m[i]++
			n /= i
		}
	}
	if n != 1 {
		m[n]++
	}
	return m
}
func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}

func union(a, b map[int]int) map[int]int {
	r := map[int]int{}
	for k, v := range a {
		r[k] += v
	}
	for k, v := range b {
		r[k] += v
	}

	return r
}

const mod = 1_000_000_000 + 7

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	if n == 1 {
		fmt.Fprintf(writer, "%v\n", 1)
		return
	}
	a := make([]map[int]int, n-1)
	for i := 2; i <= n; i++ {
		a[i-2] = primeFactor(i)
	}

	m := a[0]
	for i := 1; i < n-1; i++ {
		m = union(m, a[i])
	}
	ans := 1
	for _, v := range m {

		ans *= (v + 1)
		ans %= mod
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
