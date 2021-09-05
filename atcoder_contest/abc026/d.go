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
	var a float64
	fmt.Fscan(reader, &a)
	var b float64
	fmt.Fscan(reader, &b)
	var c float64
	fmt.Fscan(reader, &c)

	a /= 100.0
	b /= 100.0

	low := 0.0
	high := 200.0
	for i := 0; i < 45; i++ {
		mid := (high + low) / 2.0
		if a*mid+b*math.Sin(c*mid*math.Pi) <= 1.0 {
			low = mid
		} else {
			high = mid
		}
	}
    fmt.Fprintf(writer, "%v\n", low)
}
