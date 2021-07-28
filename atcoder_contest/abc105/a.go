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

	if n % k == 0 {
		fmt.Fprintf(writer, "%v\n", 0)
	} else {
		fmt.Fprintf(writer, "%v\n", 1)
	}

}
