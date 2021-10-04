package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}
func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(reader, &a)
		var b int
		fmt.Fscan(reader, &b)

		if a == b {
			fmt.Fprintf(writer, "%v\n", -1)
		} else {
			fmt.Fprintf(writer, "%v\n", abs(a-b))
		}
	}
}

