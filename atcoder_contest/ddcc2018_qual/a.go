package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func pow(a, x int) int {
	r := 1
	for x > 0 {
		if x&1 == 1 {
			r *= a
		}
		a *= a
		x >>= 1
	}
	return r
}
func main() {
	defer writer.Flush()

	var x int
	fmt.Fscan(reader, &x)
	fmt.Fprintf(writer, "%v\n", pow(4, x))
}
