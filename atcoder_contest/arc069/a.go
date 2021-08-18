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
	var m int
	fmt.Fscan(reader, &m)

	if n*2 > m {
		fmt.Fprintf(writer, "%v\n", m/2)
		return
	}

	fmt.Fprintf(writer, "%v\n", n+(m-2*n)/4)
}
