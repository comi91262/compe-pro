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

	if 30 <= x {
		fmt.Fprintf(writer, "%v\n", "Yes")
	} else {
		fmt.Fprintf(writer, "%v\n", "No")
	}

}
