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

func main() {
	defer writer.Flush()

	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var w int
	fmt.Fscan(reader, &w)
	w *= 1000

	r := (w + a - 1) / a
	l := w / b

	mx := 0
	mn := 1 << 60
	for i := l; i <= r; i++ {
		if a*i <= w && w <= b*i {
			mn = min(mn, i)
			mx = max(mx, i)
		}
	}

	if mx == 0 {
		fmt.Fprintf(writer, "%v\n", "UNSATISFIABLE")
	} else {
		fmt.Fprintf(writer, "%v %v\n", mn, mx)
	}
}
