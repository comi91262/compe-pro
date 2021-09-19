package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func toDigits(x, base int) []int {
	if x == 0 {
		return []int{0}
	}

	r := make([]int, 0)
	for x != 0 {
		r = append(r, x%base)
		x = x / base
	}

	return r
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	d := toDigits(n+1, 10)

	if d[len(d)-1] == 1 {
		fmt.Fprintf(writer, "%v\n", (len(d)-1)*9)
		return
	}
	fmt.Fprintf(writer, "%v\n", (len(d)-1)*9+d[len(d)-1]-1)
}
