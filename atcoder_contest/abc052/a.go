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
	var c int
	fmt.Fscan(reader, &c)
	var d int
	fmt.Fscan(reader, &d)

	if a*b >= c*d {
		fmt.Fprintf(writer, "%v\n", a*b)
	} else {

		fmt.Fprintf(writer, "%v\n", c*d)
	}

}