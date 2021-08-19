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
	var h1 int
	fmt.Fscan(reader, &h1)
	var m1 int
	fmt.Fscan(reader, &m1)
	var h2 int
	fmt.Fscan(reader, &h2)
	var m2 int
	fmt.Fscan(reader, &m2)

	var k int
	fmt.Fscan(reader, &k)

	fmt.Fprintf(writer, "%v\n", h2*60+m2-h1*60-m1-k)
}
