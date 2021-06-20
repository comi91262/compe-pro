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

	var n int
	fmt.Fscan(reader, &n)

	a := int(float64(n) * 1.08)

	if a > 206 {
		fmt.Fprintf(writer, ":(\n")
	} else if a < 206 {
		fmt.Fprintf(writer, "Yay!\n")
	} else {
		fmt.Fprintf(writer, "so-so\n")
	}
}
