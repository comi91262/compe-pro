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

	var x, y int
	fmt.Fscan(reader, &x, &y)

	if x == y {
		fmt.Fprintf(writer, "%d\n", x)
		return
	}

	if x == 0 && y == 1 || x == 1 && y == 0 {
		fmt.Fprintf(writer, "%d\n", 2)
		return
	}

	if x == 0 && y == 2 || x == 2 && y == 0 {
		fmt.Fprintf(writer, "%d\n", 1)
		return
	}

	if x == 1 && y == 2 || x == 2 && y == 1 {
		fmt.Fprintf(writer, "%d\n", 0)
		return
	}
}
