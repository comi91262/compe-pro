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

	var a = make([]int, n)
	var p = make([]int, n)
	var x = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		fmt.Fscan(reader, &p[i])
		fmt.Fscan(reader, &x[i])
	}

	ans := 1 << 60
	for i := 0; i < n; i++ {
		if x[i]-a[i] > 0 {
			ans = min(ans, p[i])
		}
	}
	if ans == 1<<60 {
		fmt.Fprintf(writer, "%v\n", -1)
	} else {
		fmt.Fprintf(writer, "%v\n", ans)
	}
}
