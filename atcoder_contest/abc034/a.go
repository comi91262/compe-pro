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

	if x > y {
		fmt.Fprintf(writer, "%v\n", "Worse")
	} else {
		fmt.Fprintf(writer, "%v\n", "Better")
	}
}
