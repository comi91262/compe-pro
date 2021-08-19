package main

import (
	"bufio"
	"fmt"
	"os"
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

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var a, b int

	mx := 0

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a)
		fmt.Fscan(reader, &b)

		mx = max(mx, a+b)
	}
	fmt.Fprintf(writer, "%v\n", mx)
}
