package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func f(y, m, d int) int {
	return 365*y + y/4 - y/100 + y/400 + 306*(m+1)/10 + d - 429
}

func main() {
	defer writer.Flush()

	var y int
	fmt.Fscan(reader, &y)
	var m int
	fmt.Fscan(reader, &m)
	var d int
	fmt.Fscan(reader, &d)

	if m == 1 || m == 2 {
		y--
		m += 12
	}

	fmt.Fprintf(writer, "%v\n", f(2014, 5, 17)-f(y, m, d))
}
