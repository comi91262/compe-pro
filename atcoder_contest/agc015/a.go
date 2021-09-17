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
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

	if a > b {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}

	if n == 1 {
		if a != b {
			fmt.Fprintf(writer, "%v\n", 0)
			return
		}
		fmt.Fprintf(writer, "%v\n", 1)
		return
	}

	fmt.Fprintf(writer, "%v\n", (b-a)*(n-2)+1)
}
