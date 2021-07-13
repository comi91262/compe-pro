package main

import (
	"bufio"
	"fmt"
	"os"
)

var a [2000][2000]int
var b [2000][2000]int
var row, col [2000]int

func main() {
	r := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var h, w int
	fmt.Fscan(r, &h, &w)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fscan(r, &a[i][j])
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			row[i] += a[i][j]
			col[j] += a[i][j]
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			b[i][j] = row[i] + col[j] - a[i][j]
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fprint(writer, b[i][j])
			if j != w-1 {
				fmt.Fprint(writer, " ")
			}

		}
		if i != h-1 {
			fmt.Fprintln(writer, "")
		}
	}

}
