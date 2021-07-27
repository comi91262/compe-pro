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

	if a == 1 {
		a = 14
	}
	if b == 1 {
		b = 14
	}

	if a > b {
		fmt.Fprintf(writer, "%v\n", "Alice")

	} else if b > a {
		fmt.Fprintf(writer, "%v\n", "Bob")

	} else {
		fmt.Fprintf(writer, "%v\n", "Draw")
	}

}
