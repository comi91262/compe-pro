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

	i := 1
	for i*(i+1)/2 < x {
		i++
	}

	fmt.Fprintf(writer, "%v\n", i)
}
