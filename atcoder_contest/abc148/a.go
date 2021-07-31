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

	if a*b == 2 {
		fmt.Fprintf(writer, "%v\n", 3)
	} else if a*b == 3 {
		fmt.Fprintf(writer, "%v\n", 2)
	} else {
		fmt.Fprintf(writer, "%v\n", 1)
	}

}
