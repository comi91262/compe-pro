package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func digit(x int) int {
	r := 0
	for x > 0 {
		r++
		x /= 10
	}

	return r
}
func main() {
	defer writer.Flush()

	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

	fmt.Fprintf(writer, "%v\n", digit(a)*digit(b))

}
