package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var h, w int
	fmt.Fscan(reader, &h, &w)

	if h == 1 {
		fmt.Fprintf(writer, "%d\n", w)
		return
	}

	if w == 1 {
		fmt.Fprintf(writer, "%d\n", h)
		return
	}

	p := 0
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if i%2 == 0 && j%2 == 0 {
				p++
			}
		}
	}

	fmt.Fprintf(writer, "%d\n", p)
}
