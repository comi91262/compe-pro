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

	var a int
	fmt.Fscan(reader, &a)
	var p int
	fmt.Fscan(reader, &p)

	sum := a*3 + p
	fmt.Fprintf(writer, "%v\n", sum/2)
}
