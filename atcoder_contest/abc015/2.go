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
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	m := 0
	c := 0
	for i := 0; i < n; i++ {
		if a[i] == 0 {
			continue
		}
		m += 1
		c += a[i]
	}

	fmt.Fprintf(writer, "%v\n", (c+m-1)/m)

}
