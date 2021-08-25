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

	var a float64
	fmt.Fscan(reader, &a)
	var b float64
	fmt.Fscan(reader, &b)

	var h float64
	fmt.Fscan(reader, &h)
	var m float64
	fmt.Fscan(reader, &m)

	theta := 6.0 * m
	tau := math.Mod(0.5*(60.0*h+m), 360.0)

	deg := math.Abs(theta - tau)

	ans := math.Sqrt(a*a + b*b - 2*a*b*math.Cos(deg/180*math.Pi))
	fmt.Fprintf(writer, "%v\n", ans)
}
