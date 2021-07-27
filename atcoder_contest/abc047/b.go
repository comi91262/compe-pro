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

	var w int
	fmt.Fscan(reader, &w)
	var h int
	fmt.Fscan(reader, &h)
	var n int
	fmt.Fscan(reader, &n)

	var x = make([]int, n)
	var y = make([]int, n)
	var a = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
		fmt.Fscan(reader, &y[i])
		fmt.Fscan(reader, &a[i])
	}

	sx, sy, tx, ty := 0, 0, w, h

	for i := 0; i < n; i++ {
		switch a[i] {
		case 1:
			if sx <= x[i] {
				sx = x[i]
			}
		case 2:
			if x[i] <= tx {
				tx = x[i]
			}
		case 3:
			if sy <= y[i] {
				sy = y[i]
			}
		case 4:
			if y[i] <= ty {
				ty = y[i]
			}
		}
	}

	s := (tx - sx) * (ty - sy)
	if tx >= sx && ty >= sy {
		fmt.Fprintf(writer, "%v\n", s)
	} else {
		fmt.Fprintf(writer, "%v\n", 0)
	}
}
