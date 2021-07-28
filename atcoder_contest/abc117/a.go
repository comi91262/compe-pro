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

	var t float64
	fmt.Fscan(reader, &t)
	var x float64
	fmt.Fscan(reader, &x)

	fmt.Fprintf(writer, "%v\n", t/x)

}
