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

	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var c int
	fmt.Fscan(reader, &c)
	var x int
	fmt.Fscan(reader, &x)
	var y int
	fmt.Fscan(reader, &y)

	if x > y {
		fmt.Fprintf(writer, "%v\n", min(a*x+b*y, 2*c*y+(x-y)*a, 2*c*max(x, y)))
	} else {
		fmt.Fprintf(writer, "%v\n", min(a*x+b*y, 2*c*x+(y-x)*b, 2*c*max(x, y)))

	}
}
