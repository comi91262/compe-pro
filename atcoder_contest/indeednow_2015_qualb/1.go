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
func main() {
	defer writer.Flush()

	var x int
	fmt.Fscan(reader, &x)
	var y int
	fmt.Fscan(reader, &y)
	var s int
	fmt.Fscan(reader, &s)
	var t int
	fmt.Fscan(reader, &t)
	fmt.Fprintf(writer, "%v\n", abs(x-s)+abs(y-t)+1)
}
