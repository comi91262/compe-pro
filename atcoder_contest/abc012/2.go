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

	h := n / 3600
	n %= 3600
	m := n / 60
	n %= 60
	s := n % 60

	fmt.Fprintf(writer, "%02v:%02v:%02v\n", h, m, s)
}
