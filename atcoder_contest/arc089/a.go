package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var t = make([]int, n)
	var x = make([]int, n)
	var y = make([]int, n)

	var to, xo, yo int

	for i := 0; i < n; i++ {
		var tn, xn, yn int
		fmt.Fscan(reader, &tn, &xn, &yn)

		t[i] = tn - to
		x[i] = abs(xn - xo)
		y[i] = abs(yn - yo)

		to = tn
		xo = xn
		yo = yn
	}

	for i := 0; i < n; i++ {
		mn := x[i] + y[i]
		if t[i] >= mn && (t[i]-mn)%2 == 0 {
			continue
		}

		fmt.Fprintf(writer, "%v\n", "No")
		return
	}

	fmt.Fprintf(writer, "%v\n", "Yes")
}
