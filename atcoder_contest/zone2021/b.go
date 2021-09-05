package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type pair struct {
	first  float64
	second float64
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var x float64
	fmt.Fscan(reader, &x)
	var y float64
	fmt.Fscan(reader, &y)

	var d = make([]float64, n)
	var h = make([]float64, n)
	var p = make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &d[i])
		fmt.Fscan(reader, &h[i])
		p[i].first = d[i]
		p[i].second = h[i]
	}

	ans := 0.0

	for i := 0; i < n; i++ {
		ans = math.Max(ans, y-(y-p[i].second)/(x-p[i].first)*x)
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
