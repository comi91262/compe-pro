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
	for i := 1; i <= n; i++ {
		ans += i * (n / i * (n/i + 1) / 2)
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
