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
	var n int
	fmt.Fscan(reader, &n)
	var h int
	fmt.Fscan(reader, &h)
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)
	var c int
	fmt.Fscan(reader, &c)
	var d int
	fmt.Fscan(reader, &d)
	var e int
	fmt.Fscan(reader, &e)

	ans := 1 << 60
	for x := 0; x <= n; x++ {
		y := max(int(float64((n-x)*e-h-b*x)/float64(d+e)+1.0), 0)
		if x+y <= n {
			ans = min(ans, a*x+c*y)
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
