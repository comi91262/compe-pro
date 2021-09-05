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

	ans := 0.0
	for i := 1; i < n; i++ {
		ans += (float64(n) / float64(n-i))
	}
	fmt.Fprintf(writer, "%v\n", ans)
}

