package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}
func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func calc(h, w, mn int) int {
	for i := 1; i <= w-2; i++ {
		a := h * i
		b := h * ((w - i + 1) / 2)
		c := h * ((w - i) / 2)
		mn = min(mn, max(a, b, c)-min(a, b, c))
	}

	for i := 1; i <= w-1; i++ {
		a := h * i
		b := (h / 2) * (w - i)
		c := (h + 1) / 2 * (w - i)
		mn = min(mn, max(a, b, c)-min(a, b, c))
	}
	return mn
}

func main() {
	defer writer.Flush()

	var h int
	fmt.Fscan(reader, &h)
	var w int
	fmt.Fscan(reader, &w)

	fmt.Fprintf(writer, "%v\n", min(calc(h, w, 1<<60), calc(w, h, 1<<60)))
}
