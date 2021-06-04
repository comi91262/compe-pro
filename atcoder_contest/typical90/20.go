package main

import (
	"bufio"
	"fmt"
	"os"
)

func pow(a, x int) int {
	r := 1
	for x > 0 {
		r *= a
		x--
	}
	return r
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b, c int
	fmt.Fscan(reader, &a, &b, &c)

	l := a
	r := pow(c, b)

	if l < r {
		fmt.Fprintf(writer, "Yes\n")
	} else {
		fmt.Fprintf(writer, "No\n")
	}
}

