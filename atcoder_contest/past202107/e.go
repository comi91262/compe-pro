package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func pow(a, x int) int {
	r := 1
	for x > 0 {
		if x&1 == 1 {
			r *= a
		}
		a *= a
		x >>= 1
	}
	return r
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	if pow(3, 30)+1 > n {
		fmt.Fprintf(writer, "%v\n", -1)
		return
	}

	ans := n - pow(3, 30)

	k := 1
	for k < 31 {
		if ans == pow(3, 30-k) {
			fmt.Fprintf(writer, "%v\n", k)
			return
		}
		k++
	}

	fmt.Fprintf(writer, "%v\n", -1)
}
