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

	var n int
	fmt.Fscan(reader, &n)

	switch n {
	case 1:
		fmt.Fprintf(writer, "%v\n", "Hello World")
	case 2:
		var a int
		fmt.Fscan(reader, &a)
		var b int
		fmt.Fscan(reader, &b)
		fmt.Fprintf(writer, "%v\n", a+b)
	}

}
