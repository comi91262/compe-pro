package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

	var x = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
	}

	ans := 0
	s := x[0]
	for i := 0; i < n; i++ {
		ans += min((x[i]-s)*a, b)
		s = x[i]
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
