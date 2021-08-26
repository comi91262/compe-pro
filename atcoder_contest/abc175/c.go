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
	var x int
	fmt.Fscan(reader, &x)
	var k int
	fmt.Fscan(reader, &k)
	var d int
	fmt.Fscan(reader, &d)

	cnt := abs(x / d)
	if k < cnt {
		fmt.Fprintf(writer, "%v\n", abs(x)-k*d)
		return
	}

	mn := abs(x % d)
	k -= cnt

	if k%2 == 0 {
		fmt.Fprintf(writer, "%v\n", mn)
	} else {
		fmt.Fprintf(writer, "%v\n", abs(d-mn))
	}
}
