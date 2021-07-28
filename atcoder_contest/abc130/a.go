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
	var a int
	fmt.Fscan(reader, &a)

	if x < a {
		fmt.Fprintf(writer, "%v\n", 0)
	} else {
		fmt.Fprintf(writer, "%v\n", 10)
	}

}
