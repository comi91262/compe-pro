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
	var w int
	fmt.Fscan(reader, &w)

	if h == 1 || w == 1 {
		fmt.Fprintf(writer, "%v\n", 1)
		return
	}
	fmt.Fprintf(writer, "%v\n", (h*w+1)/2)
}
