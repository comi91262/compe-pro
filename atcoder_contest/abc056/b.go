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

	var w int
	fmt.Fscan(reader, &w)
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

	if b+w < a {
		fmt.Fprintf(writer, "%v\n", a-b-w)
	} else if a+w < b {
		fmt.Fprintf(writer, "%v\n", b-a-w)
	} else {
		fmt.Fprintf(writer, "%v\n", 0)
	}

}
