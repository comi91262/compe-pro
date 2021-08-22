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

	var k int
	fmt.Fscan(reader, &k)
	var n int
	fmt.Fscan(reader, &n)

	for i := 0; i < 2*(k-1)+1; i++ {
		fmt.Fprintf(writer, "%v", n-k+1+i)
		if i != 2*k {
			fmt.Fprintf(writer, " ")
		}
	}
	fmt.Fprintf(writer, "\n")
}
