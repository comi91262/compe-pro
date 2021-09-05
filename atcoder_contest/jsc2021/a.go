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
	var y int
	fmt.Fscan(reader, &y)
	var z int
	fmt.Fscan(reader, &z)

	fmt.Fprintf(writer, "%v\n", (y*z+x-1)/x-1)
}

