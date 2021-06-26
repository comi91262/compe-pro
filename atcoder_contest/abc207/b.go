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

	var a, b, c, d int
	fmt.Fscan(reader, &a, &b, &c, &d)

	m := a
	n := 0

	if b >= d*c {
		fmt.Fprintf(writer, "%v\n", -1)
		return
	}

	for i := 1; ; i++ {
		m += b
		n += c
		if m <= n*d {
			fmt.Fprintf(writer, "%v\n", i)
			break
		}
	}

}
