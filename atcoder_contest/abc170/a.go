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
	var e int
	fmt.Fscan(reader, &e)

	if a == 0 {
		fmt.Fprintf(writer, "%v\n", 1)
		return
	}
	if b == 0 {
		fmt.Fprintf(writer, "%v\n", 2)
		return
	}
	if c == 0 {
		fmt.Fprintf(writer, "%v\n", 3)
		return
	}
	if d == 0 {
		fmt.Fprintf(writer, "%v\n", 4)
		return
	}
	if e == 0 {
		fmt.Fprintf(writer, "%v\n", 5)
		return
	}

}
