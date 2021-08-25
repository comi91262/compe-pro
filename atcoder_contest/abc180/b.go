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
	var n int
	fmt.Fscan(reader, &n)
	var a = make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	x := 0.0
	for i := 0; i < n; i++ {
		x += math.Abs(a[i])
	}
	fmt.Fprintf(writer, "%v\n", x)

	z := 0.0
	for i := 0; i < n; i++ {
		z += (a[i] * a[i])
	}
	fmt.Fprintf(writer, "%v\n", math.Sqrt(z))

	y := 0.0
	for i := 0; i < n; i++ {
		y = math.Max(y, math.Abs(a[i]))
	}
	fmt.Fprintf(writer, "%v\n", y)
}
