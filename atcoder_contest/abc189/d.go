package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var s = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
	}

	y := make([][]int, 2)
	for i := 0; i < 2; i++ {
		y[i] = make([]int, n+1)
	}
	y[1][0] = 1
	y[0][0] = 1
	for i := 0; i < n; i++ {
		switch s[i] {
		case "OR":
			y[1][i+1] += 2*y[1][i] + y[0][i]
			y[0][i+1] += y[0][i]
		case "AND":
			y[1][i+1] += y[1][i]
			y[0][i+1] += 2*y[0][i] + y[1][i]
		}
	}
	fmt.Fprintf(writer, "%v\n", y[1][n])

}
