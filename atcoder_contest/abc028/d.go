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

	var n int
	fmt.Fscan(reader, &n)
	var k int
	fmt.Fscan(reader, &k)

	// 3
	a := (k - 1 - 1 + 1) * (n - k - 1 + 1) * 6
	b := (n - k - 1 + 1) * 3
	c := (k - 1 - 1 + 1) * 3
	d := 1

	fmt.Fprintf(writer, "%v\n", float64(a+b+c+d)/float64(n*n*n))
}
