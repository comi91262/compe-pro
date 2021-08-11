package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func pow(a, x int) int {
	r := 1
	for x > 0 {
		if x&1 == 1 {
			r *= a
		}
		a *= a
		x >>= 1
	}
	return r
}
func max(arg ...int) int {
	max := arg[0]
	for _, x := range arg {
		if max < x {
			max = x
		}
	}
	return max
}

func main() {
	defer writer.Flush()

	var x int
	fmt.Fscan(reader, &x)

	lim := int(math.Sqrt(float64(x)))

	ans := 0
	for i := 1; i <= lim; i++ {
		j := 2
		for i != 1 && x >= pow(i, j) {
			// fmt.Fprintf(writer, "%v %v %v\n", i, j, pow(i, j))
			j++
		}
		if x == pow(i, j) {
			ans = max(ans, pow(i, j))
		} else {
			ans = max(ans, pow(i, j-1))
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
