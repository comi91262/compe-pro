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

	var h1, h2 int
	fmt.Fscan(reader, &h1, &h2)

	fmt.Fprintf(writer, "%v\n", h1-h2)

}
