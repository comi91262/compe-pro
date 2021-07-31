package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const mod = 1_000_000_000 + 7

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	ans := 1
	for i := 1; i <= n; i++ {
		ans *= i
		ans %= mod
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
