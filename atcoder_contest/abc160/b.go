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

	var x int
	fmt.Fscan(reader, &x)
	ans := 0
	ans += x / 500 * 1000
	ans += (x % 500) / 5 * 5
	fmt.Fprintf(writer, "%v\n", ans)
}
