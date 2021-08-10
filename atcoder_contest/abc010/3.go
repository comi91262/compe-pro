package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

func main() {
	defer writer.Flush()

	var txa int
	fmt.Fscan(reader, &txa)
	var tya int
	fmt.Fscan(reader, &tya)
	var txb int
	fmt.Fscan(reader, &txb)
	var tyb int
	fmt.Fscan(reader, &tyb)

	var t int
	fmt.Fscan(reader, &t)
	var v int
	fmt.Fscan(reader, &v)
	var n int
	fmt.Fscan(reader, &n)
	var x = make([]int, n)
	var y = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
		fmt.Fscan(reader, &y[i])
	}

	for i := 0; i < n; i++ {
		xa := abs(x[i] - txa)
		ya := abs(y[i] - tya)

		xb := abs(x[i] - txb)
		yb := abs(y[i] - tyb)

		td := math.Sqrt(float64(xa*xa + ya*ya))
		sd := math.Sqrt(float64(xb*xb + yb*yb))
		d := float64(v * t)

		if d >= td+sd {
			fmt.Fprintf(writer, "%v\n", "YES")
			return
		}

	}

	fmt.Fprintf(writer, "%v\n", "NO")
}
