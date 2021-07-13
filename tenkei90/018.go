package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var t, l, x, y, q int
	fmt.Fscan(reader, &t, &l, &x, &y, &q)

	var e = make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Fscan(reader, &e[i])
	}

	for i := 0; i < q; i++ {
		theta := 2.0*math.Pi*float64(e[i])/float64(t) - math.Pi/2.0
		r := float64(l) / 2.0
		a := -r * math.Cos(theta)
		b := r + r*math.Sin(theta)

		fx := float64(x)
		fy := float64(y)

		u := math.Sqrt(math.Pow(fx, 2.0) + math.Pow(a-fy, 2.0))

		fmt.Fprintf(writer, "%.9f\n", 90.0-math.Atan2(u, b)/math.Pi*180.0)

	}

}
