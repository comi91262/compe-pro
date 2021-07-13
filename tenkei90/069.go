package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const d = 1_000_000_000 + 7

func pow(a, x int) int {
	r := 1
	for x > 0 {
		if x&1 == 1 {
			r = r * a % d
		}
		a = a * a % d
		x >>= 1
	}
	return r
}

func main() {
	defer writer.Flush()

	var n, k int
	fmt.Fscan(reader, &n, &k)

	if n == 1 {
		fmt.Fprintf(writer, "%d\n", k)
	} else if n == 2 {
		fmt.Fprintf(writer, "%d\n", k*(k-1)%d)
	} else {
		if k <= 2 {
			fmt.Fprintf(writer, "%d\n", 0)
		} else {
			pre := k * (k - 1) % d
			fmt.Fprintf(writer, "%d\n", pre*pow(k-2, n-2)%d)
		}
	}

}
