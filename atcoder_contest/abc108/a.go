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

	var k int
	fmt.Fscan(reader, &k)

  fmt.Fprintf(writer, "%v\n", k/2*((k+1)/2))
}
