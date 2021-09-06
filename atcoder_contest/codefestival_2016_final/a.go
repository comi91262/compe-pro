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

	s := make([][]string, h)
	for i := 0; i < h; i++ {
		s[i] = make([]string, w)
		for j := 0; j < w; j++ {
			fmt.Fscan(reader, &s[i][j])
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == "snuke" {
				fmt.Fprintf(writer, "%v%v\n", string("A"[0]+byte(j)), i+1)
				return
			}
		}
	}
}
