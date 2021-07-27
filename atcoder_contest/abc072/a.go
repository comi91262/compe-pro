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

	var x int
	fmt.Fscan(reader, &x)
	var t int
	fmt.Fscan(reader, &t)

	fmt.Fprintf(writer, "%v\n", max(x-t, 0))
}
