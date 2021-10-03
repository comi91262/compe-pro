package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var t int
	fmt.Fscan(reader, &t)
	var n int
	fmt.Fscan(reader, &n)

	x := (100*n + t - 1) / t
	fmt.Fprintf(writer, "%v\n", x+n-1)
}
