package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

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

	var n int
	fmt.Fscan(reader, &n)

	mn := 1 << 60
	for i := 1; i*i <= n; i++ {
		for j := 1; j*i <= n; j++ {
			// fmt.Fprintf(writer, "%v %v\n", i, j)
			mn = min(abs(i-j)+n-i*j, mn)
		}
	}
	fmt.Fprintf(writer, "%v\n", mn)
}
