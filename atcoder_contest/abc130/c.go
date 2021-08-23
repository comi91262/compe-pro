package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

const sig = 0.000001

func main() {
	defer writer.Flush()
	var w float64
	fmt.Fscan(reader, &w)
	var h float64
	fmt.Fscan(reader, &h)

	var x float64
	fmt.Fscan(reader, &x)
	var y float64
	fmt.Fscan(reader, &y)

	if math.Abs(w/2.0-x) < sig && math.Abs(h/2.0-y) < sig {
		fmt.Fprintf(writer, "%v %v\n", w*h/2, 1)
		return
	}

	if x == 0 || x == w {
		fmt.Fprintf(writer, "%v %v\n", w*h/2, 0)
		return
	}

	if y == 0 || y == h {
		fmt.Fprintf(writer, "%v %v\n", w*h/2, 0)
		return
	}

	fmt.Fprintf(writer, "%v %v\n", w*h/2, 0)
	return
}
