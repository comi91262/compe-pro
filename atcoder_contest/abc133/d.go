package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func sum(a []int) int {
	r := 0
	for i := range a {
		r += a[i]
	}
	return r
}
func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
		a[i] *= 2
	}

	x := make([]int, n)
	total := sum(a) / 2
	for i := 1; i < n; i += 2 {
		total -= a[i]
	}
	x[0] = total

	for i := 0; i < n-1; i++ {
		x[i+1] = a[i] - x[i]
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%v", x[i])
		if i != n-1 {
			fmt.Fprintf(writer, " ")
		}
	}
	fmt.Fprintf(writer, "\n")

}
