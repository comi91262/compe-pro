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

	var n, x int
	fmt.Fscan(reader, &n, &x)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	sum := 0
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			sum += a[i] - 1
		} else {
			sum += a[i]
		}
	}

	if sum <= x {
		fmt.Fprintf(writer, "%v\n", "Yes")
	} else {
		fmt.Fprintf(writer, "%v\n", "No")
	}
}
