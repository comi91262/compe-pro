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

	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

	an := 0
	bn := 0

	for a > 0 {
		an += a % 10
		bn += b % 10

		a /= 10
		b /= 10
	}

	fmt.Fprintf(writer, "%v\n", max(an, bn))
}
