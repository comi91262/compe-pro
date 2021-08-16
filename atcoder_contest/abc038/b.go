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
	var h1 int
	fmt.Fscan(reader, &h1)
	var w1 int
	fmt.Fscan(reader, &w1)
	var h2 int
	fmt.Fscan(reader, &h2)
	var w2 int
	fmt.Fscan(reader, &w2)

	if h1 == h2 {
		fmt.Fprintf(writer, "%v\n", "YES")
		return
	} else if h1 == w2 {
		fmt.Fprintf(writer, "%v\n", "YES")
		return
	} else if w1 == h2 {
		fmt.Fprintf(writer, "%v\n", "YES")
		return
	} else if w1 == w2 {
		fmt.Fprintf(writer, "%v\n", "YES")
		return
	}

	fmt.Fprintf(writer, "%v\n", "NO")
}
