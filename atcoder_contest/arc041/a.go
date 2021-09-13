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

	var x int
	fmt.Fscan(reader, &x)
	var y int
	fmt.Fscan(reader, &y)
	var k int
	fmt.Fscan(reader, &k)

	if y >= k {
		fmt.Fprintf(writer, "%v\n", x+k)
	} else {
		fmt.Fprintf(writer, "%v\n", x-(k-y)+y)
	}

}
