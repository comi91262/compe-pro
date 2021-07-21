package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	var l = make([]int, n)
	for i := 1; i < n; i++ {
		l[i] = gcd(l[i-1], a[i-1])
	}

	var r = make([]int, n)
	for i := n - 2; i >= 0; i-- {
		r[i] = gcd(r[i+1], a[i+1])
	}

	ans := 0
	for i := 0; i < n; i++ {
		g := gcd(l[i], r[i])
		ans = max(ans, g)
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
