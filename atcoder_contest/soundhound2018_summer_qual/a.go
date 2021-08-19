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
	var b int
	fmt.Fscan(reader, &b)
	if a+b == 15 {
		fmt.Fprintf(writer, "%v\n", "+")
		return
	}
	if a*b == 15 {
		fmt.Fprintf(writer, "%v\n", "*")
		return
	}
	fmt.Fprintf(writer, "%v\n", "x")
}
