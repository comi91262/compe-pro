package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var l float64
	fmt.Fscan(reader, &l)
	var x float64
	fmt.Fscan(reader, &x)
	var y float64
	fmt.Fscan(reader, &y)
	var s float64
	fmt.Fscan(reader, &s)
	var d float64
	fmt.Fscan(reader, &d)

	var ans float64
	if s <= d {
		ans = math.Abs(s-d) / (x + y)
		if y > x {
			ans = math.Min(ans, math.Abs(s+l-d)/math.Abs(x-y))
		}
	} else {
		ans = math.Abs(l-s+d) / (x + y)
		if y > x {
			ans = math.Min(ans, math.Abs(s-d)/math.Abs(x-y))
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
