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

	var r float64
	fmt.Fscan(reader, &r)
	var x float64
	fmt.Fscan(reader, &x)
	var y float64
	fmt.Fscan(reader, &y)

	d := math.Sqrt(x*x + y*y)

	if d < r {
		fmt.Fprintf(writer, "%v\n", 2)
		return
	}

	ans := 0
	for float64(ans)*r < d {
		ans++
	}

	fmt.Fprintf(writer, "%v\n", ans)

}
