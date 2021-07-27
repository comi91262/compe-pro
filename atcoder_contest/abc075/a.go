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

	switch {
	case a == b:
		fmt.Fprintf(writer, "%v\n", c)
	case b == c:
		fmt.Fprintf(writer, "%v\n", a)
	case c == a:
		fmt.Fprintf(writer, "%v\n", b)
	}

}
