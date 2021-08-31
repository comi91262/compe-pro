package main

import (
	"bufio"
	"fmt"
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

func main() {
	defer writer.Flush()

	var x int
	fmt.Fscan(reader, &x)

	for a := -120; a <= 120; a++ {
		for b := -120; b <= 120; b++ {
			if pow(a, 5)-pow(b, 5) == x {
				fmt.Fprintf(writer, "%v %v\n", a, b)
				return
			}
		}
	}
}

