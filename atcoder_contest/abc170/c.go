package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func contains(x int, a []int) bool {
	for _, e := range a {
		if x == e {
			return true
		}
	}
	return false
}
func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
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

func main() {
	defer writer.Flush()

	var x int
	fmt.Fscan(reader, &x)
	var n int
	fmt.Fscan(reader, &n)
	var p = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i])
	}

	ans := 0
	mn := 1 << 60
	for i := 0; i <= 101; i++ {
		if !contains(i, p) {
			if mn > abs(x-i) {
				mn = abs(x - i)
				ans = i
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
