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

	for _, v := range toDigits(n, 10) {
		if v == 3 {
			fmt.Fprintf(writer, "%v\n", "YES")
			return
		}
	}

	if n%3 == 0 {
		fmt.Fprintf(writer, "%v\n", "YES")
		return
	}

	fmt.Fprintf(writer, "%v\n", "NO")
	return
}
