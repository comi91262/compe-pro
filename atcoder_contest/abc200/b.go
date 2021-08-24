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

	for i := 0; i < k; i++ {
		if n%200 == 0 {
			n /= 200
		} else {
			n = n*1000 + 200
		}
	}

	fmt.Fprintf(writer, "%v\n", n)
}
