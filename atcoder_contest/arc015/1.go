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
	var n float64
	fmt.Fscan(reader, &n)

	fmt.Fprintf(writer, "%v\n", (1.8*n)+32.0)
}
