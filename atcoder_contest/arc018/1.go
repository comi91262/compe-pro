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
	var h float64
	fmt.Fscan(reader, &h)
	var b float64
	fmt.Fscan(reader, &b)

	fmt.Fprintf(writer, "%v\n", (h*0.01)*(h*0.01)*b)
}
