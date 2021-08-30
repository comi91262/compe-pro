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
	var x float64
	fmt.Fscan(reader, &x)

	half := a * a * b / 2

	if x <= half {
		c := 2.0 * x / (a * b)
		fmt.Fprintf(writer, "%v\n", math.Atan2(b, c)*180/math.Pi)
	} else {
		c := 2.0*x/(a*a) - b
		fmt.Fprintf(writer, "%v\n", math.Atan2(b-c, a)*180/math.Pi)

	}
}
