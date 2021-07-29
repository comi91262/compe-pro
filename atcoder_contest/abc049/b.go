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

	var h int
	fmt.Fscan(reader, &h)
	var w int
	fmt.Fscan(reader, &w)

	c := make([][]string, w)
	for i := 0; i < w; i++ {
		c[i] = make([]string, h)
		for j := 0; j < h; j++ {
			fmt.Fscan(reader, &c[i][j])
		}
	}

	d := make([][]string, w)
	for i := 0; i < w; i++ {
		d[i] = make([]string, h*2)
		for j := 0; j < h*2; j++ {
			fmt.Fscan(reader, &d[i][j])
		}
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			d[i][2*j] = c[i][j]
			d[i][2*j+1] = c[i][j]
		}
	}

	for i := 0; i < w; i++ {
		for j := 0; j < 2*h; j++ {
			fmt.Fprintf(writer, "%v\n", d[i][j])
		}
	}
}
