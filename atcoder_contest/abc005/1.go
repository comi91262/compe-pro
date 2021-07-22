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

	var x,y int
	fmt.Fscan(reader, &x, &y)

	fmt.Fprintf(writer, "%v\n", y/x)
}
