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
	var h int
	fmt.Fscan(reader, &h)
	var m int
	fmt.Fscan(reader, &m)

	fmt.Fprintf(writer, "%v\n", 18*60-h*60-m)
}
