package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var x float64
	fmt.Fscan(reader, &x)
	var v = make([]float64, n)
	var p = make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &v[i])
		fmt.Fscan(reader, &p[i])
	}

	total := 0.0
	x *= 100
	for i := 0; i < n; i++ {
		total += v[i] * p[i]
		if total > x {
			fmt.Fprintf(writer, "%v\n", i+1)
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", -1)
}
