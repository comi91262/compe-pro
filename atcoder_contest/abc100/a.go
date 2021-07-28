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

	if 8-a >= 0 && 8-b >= 0 {
		fmt.Fprintf(writer, "%v\n", "Yay!")
	} else {
		fmt.Fprintf(writer, "%v\n", ":(")
	}
}
