package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
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

func intersect(x, y map[string]int) map[string]int {
	m := map[string]int{}

	for k, v := range x {
		if y[k] > 0 {
			m[k] = min(y[k], v)
		}
	}
	return m
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	ans := 0
	for i := 1; i < n; i++ {
		t, u := s[:i], s[i:]

		a := map[string]int{}
		b := map[string]int{}
		for j := range t {
			a[t[j]]++
		}
		for j := range u {
			b[u[j]]++
		}
		ans = max(ans, len(intersect(a, b)))
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
