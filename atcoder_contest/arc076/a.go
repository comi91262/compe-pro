package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const mod = 1_000_000_000 + 7

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

	if n > m+1 {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}

	if m > n+1 {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}

	nn := 1
	for i := 1; i <= n; i++ {
		nn *= i
		nn %= mod

	}

	mn := 1
	for i := 1; i <= m; i++ {
		mn *= i
		mn %= mod
	}

	ans := nn * mn
	ans %= mod

	if m == n {
		ans *= 2
		ans %= mod
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
