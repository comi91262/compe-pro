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

	var x1 int
	fmt.Fscan(reader, &x1)
	var y1 int
	fmt.Fscan(reader, &y1)
	var x2 int
	fmt.Fscan(reader, &x2)
	var y2 int
	fmt.Fscan(reader, &y2)

	vx := x2 - x1
	vy := y2 - y1

	vx, vy = -vy, vx

	x3 := x2 + vx
	y3 := y2 + vy

	vx, vy = -vy, vx

	x4 := x3 + vx
	y4 := y3 + vy

	fmt.Fprintf(writer, "%v %v %v %v\n", x3, y3, x4, y4)
}
