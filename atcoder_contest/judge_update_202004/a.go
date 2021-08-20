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

	var s int
	fmt.Fscan(reader, &s)
	var l int
	fmt.Fscan(reader, &l)
	var r int
	fmt.Fscan(reader, &r)

	x := s
	for l > x || x > r {
		if x < l {
			x++
			continue
		}

		if r < x {
			x--
			continue
		}
	}

	fmt.Fprintf(writer, "%v\n", x)
}
