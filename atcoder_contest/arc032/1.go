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

	p := n * (n + 1) / 2

	if p == 1 {
		fmt.Fprintf(writer, "%v\n", "BOWWOW")
		return
	}

	for i := 2; i*i < p; i++ {
		if p%i == 0 {
			fmt.Fprintf(writer, "%v\n", "BOWWOW")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "WANWAN")
}
