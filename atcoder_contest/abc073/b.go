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

	ans := 0
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Fscan(reader, &a)
		fmt.Fscan(reader, &b)
		ans += b - a + 1
	}

	fmt.Fprintf(writer, "%v\n", ans)
}

