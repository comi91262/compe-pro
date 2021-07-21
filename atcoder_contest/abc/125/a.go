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

	var a, b, t int
	fmt.Fscan(reader, &a, &b, &t)

	if t < a {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}

	fmt.Fprintf(writer, "%v\n", t/a*b)

}
