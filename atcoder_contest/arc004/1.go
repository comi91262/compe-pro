package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type point struct {
	x float64
	y float64
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var p = make([]point, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i].x, &p[i].y)
	}

	ans := 0.0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := math.Abs(p[i].x - p[j].x)
			dy := math.Abs(p[i].y - p[j].y)
			ans = math.Max(ans, math.Sqrt(dx*dx+dy*dy))
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
