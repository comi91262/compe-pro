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

	ans := n / 10 * 100
	n %= 10

	if n >= 7 {
		ans += 100
	} else {
		ans += n * 15
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
