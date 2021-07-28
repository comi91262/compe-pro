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
	var k int
	fmt.Fscan(reader, &k)
	var s string
	fmt.Fscan(reader, &s)

	for i := 0; i < n; i++ {
		if i == k-1 {
			fmt.Fprintf(writer, "%v", string(s[i]+32))
		} else {
			fmt.Fprintf(writer, "%v", string(s[i]))
		}

	}

	fmt.Fprintf(writer, "\n")
}
