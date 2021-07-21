package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

func main() {
	defer writer.Flush()

	var xa, ya int
	fmt.Fscan(reader, &xa, &ya)
	var xb, yb int
	fmt.Fscan(reader, &xb, &yb)
	var xc, yc int
	fmt.Fscan(reader, &xc, &yc)

	a := xb - xa
	b := yb - ya

	c := xc - xa
	d := yc - ya

	ans := float64(abs(a*d-b*c)) / 2.0

	fmt.Fprintf(writer, "%v\n", ans)
}
