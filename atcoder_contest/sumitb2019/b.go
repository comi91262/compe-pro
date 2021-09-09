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

	x1 := (n + 107) / 108 * 100
	x2 := n / 108 * 100

	for i := x2; i <= x1; i++ {
		if i*108/100 == n {
			fmt.Fprintf(writer, "%v\n", i)
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", ":(")
}
