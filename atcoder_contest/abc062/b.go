package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var h int
	fmt.Fscan(reader, &h)
	var w int
	fmt.Fscan(reader, &w)

	a := make([][]string, h)
	for i := 0; i < h; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		a[i] = strings.Split(tmp, "")
	}

	var b = make([][]string, h+2)
	for i := 0; i < h+2; i++ {
		b[i] = make([]string, w+2)
	}

	for i := range b {
		for j := range b[i] {
			if i == 0 || j == 0 || i == h+1 || j == w+1 {
				b[i][j] = "#"
			} else {
				b[i][j] = a[i-1][j-1]
			}
		}
	}

	for i := 0; i < h+2; i++ {
		for j := 0; j < w+2; j++ {
			fmt.Fprintf(writer, "%v", b[i][j])
		}
		fmt.Fprintf(writer, "\n")
	}
}


