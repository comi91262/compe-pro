package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func rotate(x, y, deg float64) (float64, float64) {
	cs := math.Cos(deg * math.Pi / 180.0)
	sn := math.Sin(deg * math.Pi / 180.0)

	return x*cs - y*sn, x*sn + y*cs
}

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

	var x float64
	fmt.Fscan(reader, &x)
	var y float64
	fmt.Fscan(reader, &y)
	var xo float64
	fmt.Fscan(reader, &xo)
	var yo float64
	fmt.Fscan(reader, &yo)

	xo = xo - x
	yo = yo - y

	xo /= 2.0
	yo /= 2.0

	m := (float64(n) / 2.0)
	deg := 180.0 / m * (m - 1)
	xx, yy := rotate(xo, yo, -deg)
	fmt.Fprintf(writer, "%v %v\n", xo+xx+x, yo+yy+y)

}
