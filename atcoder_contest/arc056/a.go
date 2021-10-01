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
	var k int
	fmt.Fscan(reader, &k)
	var l int
	fmt.Fscan(reader, &l)

	fmt.Fprintf(writer, "%v\n", min(a*(k%l)+b*(k/l), (k+l-1)/l*b))

}

