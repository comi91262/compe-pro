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

	var n int
	fmt.Fscan(reader, &n)

	a := []int{1, 2, 3, 4, 5, 6}

	m := n % 30
	for i := 0; i < m; i++ {
		a[i%5+1-1], a[i%5+2-1] = a[i%5+2-1], a[i%5+1-1]
	}

	for i := 0; i < 6; i++ {
		fmt.Fprintf(writer, "%v", a[i])
	}
	fmt.Fprintf(writer, "\n")
}
