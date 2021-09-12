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

	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var x int
	fmt.Fscan(reader, &x)
	var y int
	fmt.Fscan(reader, &y)

	if a == b {
		fmt.Fprintf(writer, "%v\n", x)
	}
	if a > b {
		if a-b == 1 {
			fmt.Fprintf(writer, "%v\n", x)
		} else {
			fmt.Fprintf(writer, "%v\n", x+(a-b-1)*min(y, 2*x))
		}
	}
	if a < b {
		fmt.Fprintf(writer, "%v\n", x+(b-a)*min(y, 2*x))
	}

}
