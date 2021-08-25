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

	var l int
	fmt.Fscan(reader, &l)

	e := float64(l) / 3.0
	fmt.Fprintf(writer, "%v\n", e*e*e)

}
