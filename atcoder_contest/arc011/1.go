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

	var m int
	fmt.Fscan(reader, &m)
	var n int
	fmt.Fscan(reader, &n)
	var N int
	fmt.Fscan(reader, &N)

	ans := N
	for N >= m {
		rest := N % m
		nex := N / m * n
		ans += nex
		N = nex + rest
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
