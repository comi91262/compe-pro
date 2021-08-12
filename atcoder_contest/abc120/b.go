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

	n := 0
	for i := min(a, b); i > 0; i-- {
		if a%i == 0 && b%i == 0 {
			n++
		}

		if k == n {
			fmt.Fprintf(writer, "%v\n", i)
			return
		}
	}

}

