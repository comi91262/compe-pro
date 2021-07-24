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

	n := 4
	c := make([][]string, n)
	for i := 0; i < n; i++ {
		c[i] = make([]string, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(reader, &c[i][j])
		}
	}

	for j := 0; j < 4; j++ {
		c[0][j], c[3][j] = c[3][j], c[0][j]
	}
	for j := 0; j < 4; j++ {
		c[1][j], c[2][j] = c[2][j], c[1][j]
	}

	for j := 0; j < 4; j++ {
		c[j][0], c[j][3] = c[j][3], c[j][0]
	}
	for j := 0; j < 4; j++ {
		c[j][1], c[j][2] = c[j][2], c[j][1]
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j != 0 {
				fmt.Fprintf(writer, " ")
			}
			fmt.Fprintf(writer, "%v", c[i][j])

		}
		fmt.Fprintf(writer, "\n")
	}
}
