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
	for a := 1; a <= n-1; a++ {
		ans += (n - 1) / a
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
