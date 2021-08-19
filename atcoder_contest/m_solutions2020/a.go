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

	fmt.Fprintf(writer, "%v\n", 8-(x-400)/200)

}
