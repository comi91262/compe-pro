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
	var x = make([]int, n)
	var y = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &x[i])
		fmt.Fscan(reader, &y[i])
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				dx1 := x[i] - x[k]
				dy1 := y[i] - y[k]

				dx2 := x[j] - x[k]
				dy2 := y[j] - y[k]

				if dx1*dy2 == dy1*dx2 {
					fmt.Fprintf(writer, "%v\n", "Yes")
					return
				}
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", "No")
}
