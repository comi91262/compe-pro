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
	return int(math.Abs(float64(x)))
}

func main() {
	defer writer.Flush()

	var a, b, c int
	fmt.Fscan(reader, &a, &b, &c)

	if c%2 == 0 {
		a = abs(a)
		b = abs(b)
	}

	if a >= 0 && b >= 0 {
		if a > b {
			fmt.Fprintf(writer, ">")
		} else if a == b {
			fmt.Fprintf(writer, "=")
		} else {
			fmt.Fprintf(writer, "<")
		}
	} else if a >= 0 && b < 0 {
		fmt.Fprintf(writer, ">")
	} else if a < 0 && b >= 0 {
		fmt.Fprintf(writer, "<")
	} else if a < 0 && b < 0 {
		if a > b {
			fmt.Fprintf(writer, ">")
		} else if a == b {
			fmt.Fprintf(writer, "=")
		} else {
			fmt.Fprintf(writer, "<")
		}
	}
}
