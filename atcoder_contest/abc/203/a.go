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

	var a, b, c int
	fmt.Fscan(reader, &a, &b, &c)

	if a == b {
		fmt.Fprintf(writer, "%d\n", c)
		return
	}

	if b == c {
		fmt.Fprintf(writer, "%d\n", a)
		return
	}

	if c == a {
		fmt.Fprintf(writer, "%d\n", b)
		return
	}

	fmt.Fprintf(writer, "%d\n", 0)
}
