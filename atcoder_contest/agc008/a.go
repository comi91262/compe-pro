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
func min(arg ...int) int {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}
func main() {
	defer writer.Flush()

	var x int
	fmt.Fscan(reader, &x)
	var y int
	fmt.Fscan(reader, &y)

	ans := 1 << 60
	if x <= y {
		ans = min(ans, y-x)
	}
	x *= -1
	if x <= y {
		ans = min(ans, y-x+1)
	}
	x *= -1
	y *= -1
	if x <= y {
		ans = min(ans, y-x+1)
	}
	x *= -1
	if x <= y {
		ans = min(ans, y-x+2)
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
