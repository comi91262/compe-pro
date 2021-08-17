package main

import (
	"bufio"
	"fmt"
	"math"
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

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var t float64
	fmt.Fscan(reader, &t)
	var a float64
	fmt.Fscan(reader, &a)

	var h = make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &h[i])
	}

	for i := 0; i < n; i++ {
		h[i] = t - h[i]*0.006
	}

	ans := 0
	mn := 1000000.0
	for i := 0; i < n; i++ {
		diff := math.Abs(h[i] - a)
		if mn > diff {
			mn = diff
			ans = i + 1
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
