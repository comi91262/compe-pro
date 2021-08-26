package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	n := len(strconv.Itoa(x))

	ans := 0
	for i := 1; i < pow(10, n/2); i++ {
		d := len(strconv.Itoa(i))
		if i*pow(10, d)+i <= x {
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
