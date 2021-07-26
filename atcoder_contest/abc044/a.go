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
	var x int
	fmt.Fscan(reader, &x)
	var y int
	fmt.Fscan(reader, &y)

	if n >= k {
		fmt.Fprintf(writer, "%v\n", k*x+(n-k)*y)
	} else {
		fmt.Fprintf(writer, "%v\n", n*x)
	}

}
