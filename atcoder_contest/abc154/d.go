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
	var k int
	fmt.Fscan(reader, &k)

	var p = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i])
	}

	e := make([]float64, n)
	for i := 0; i < n; i++ {
		for j := 1; j <= p[i]; j++ {
			e[i] += 1.0 / float64(p[i]) * float64(j)
		}
	}
	for i := 0; i < n-1; i++ {
		e[i+1] += e[i]
	}

	ans := 0.0
	for i := 0; i < n-k+1; i++ {
		mx := e[i+k-1]
		if i != 0 {
			mx -= e[i-1]
		}

		ans = math.Max(mx, ans)
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
