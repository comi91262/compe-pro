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
	a := 42
	cnt := 1
	for a*pow(2, cnt) < 130_000_000/2 {
		cnt++
	}
	fmt.Fprintf(writer, "%v\n", a*pow(2, cnt+1))

}
