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
	var r int
	fmt.Fscan(reader, &r)
	var x2 int
	fmt.Fscan(reader, &x2)
	var y2 int
	fmt.Fscan(reader, &y2)
	var x3 int
	fmt.Fscan(reader, &x3)
	var y3 int
	fmt.Fscan(reader, &y3)

	isRed := true
	isBlue := false

	if x2 <= x1-r && x1+r <= x3 && y2 <= y1-r && y1+r <= y3 {
		isRed = false
	}

	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	if dx*dx+dy*dy > r*r {
		isBlue = true
	}
	dx = abs(x3 - x1)
	dy = abs(y3 - y1)
	if dx*dx+dy*dy > r*r {
		isBlue = true
	}
	dx = abs(x3 - x1)
	dy = abs(y2 - y1)
	if dx*dx+dy*dy > r*r {
		isBlue = true
	}
	dx = abs(x2 - x1)
	dy = abs(y3 - y1)
	if dx*dx+dy*dy > r*r {
		isBlue = true
	}

	if isRed {
		fmt.Fprintf(writer, "%v\n", "YES")
	} else {
		fmt.Fprintf(writer, "%v\n", "NO")
	}
	if isBlue {
		fmt.Fprintf(writer, "%v", "YES")
	} else {
		fmt.Fprintf(writer, "%v", "NO")
	}
}
