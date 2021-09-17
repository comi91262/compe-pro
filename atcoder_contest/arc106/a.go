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

	var n int
	fmt.Fscan(reader, &n)
	a := 1
	b := 1
	upper := pow(10, 18) + 1
	for pow(3, a) < upper/3 {
		a++
	}
	for pow(5, b) < upper/5 {
		b++
	}
	for i := 1; i <= a; i++ {
		for j := 1; j <= b; j++ {
			if n == pow(3, i)+pow(5, j) {
				fmt.Fprintf(writer, "%v %v\n", i, j)
				return
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", -1)
}
