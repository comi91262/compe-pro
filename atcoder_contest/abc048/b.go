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
	var x int
	fmt.Fscan(reader, &x)

	if a == 0 {
		fmt.Fprintf(writer, "%v\n", b/x+1)
	} else {
		fmt.Fprintf(writer, "%v\n", b/x-(a-1)/x)
	}
}

