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
	var t int
	fmt.Fscan(reader, &t)
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var c int
	fmt.Fscan(reader, &c)
	var d int
	fmt.Fscan(reader, &d)

	switch {
	case a+c <= t:
		fmt.Fprintf(writer, "%v\n", b+d)
	case c <= t:
		fmt.Fprintf(writer, "%v\n", d)
	case a <= t:
		fmt.Fprintf(writer, "%v\n", b)
	default:
		fmt.Fprintf(writer, "%v\n", 0)
	}

}
