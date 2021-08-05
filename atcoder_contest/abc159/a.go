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
	var m int
	fmt.Fscan(reader, &m)

	ans := 0

	if n > 1 {
		ans += n * (n - 1) / 2
	}

	if m > 1 {
		ans += m * (m - 1) / 2
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
