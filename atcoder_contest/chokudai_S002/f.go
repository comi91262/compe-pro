package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type pair struct {
	x int
	y int
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var a, b int

	set := map[pair]struct{}{}
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a)
		fmt.Fscan(reader, &b)
		if a > b {
			a, b = b, a
		}
		set[pair{a, b}] = struct{}{}
	}
	fmt.Fprintf(writer, "%v\n", len(set))
}
