package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func isSquare(n int) bool {
	pre := int(math.Sqrt(float64(n)))

	return n == pre*pre
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	ans := 0
	for i := 1; i < n; i++ {
		y := n*n - i*i

		if i*i < y && isSquare(y) {
			ans++
		}
	}
	ans *= 2

	fmt.Fprintf(writer, "%v\n", ans)
}
