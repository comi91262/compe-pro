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
	var H int
	fmt.Fscan(reader, &H)
	var W int
	fmt.Fscan(reader, &W)
	var h int
	fmt.Fscan(reader, &h)
	var w int
	fmt.Fscan(reader, &w)

	fmt.Fprintf(writer, "%v\n", (H-h)*(W-w))
}
