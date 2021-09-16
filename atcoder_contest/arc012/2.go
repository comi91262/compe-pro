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
	var va float64
	fmt.Fscan(reader, &va)
	var vb float64
	fmt.Fscan(reader, &vb)
	var l float64
	fmt.Fscan(reader, &l)

	a := 0.0
	b := l
	for i := 0; i < n; i++ {
		a, b = b, b+(b-a)/va*vb
	}
	fmt.Fprintf(writer, "%v\n", math.Abs(a-b))
}
