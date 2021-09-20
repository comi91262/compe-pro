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

	a := make([]int, 20)
	a[0] = 100
	a[1] = 100
	a[2] = 200
	for i := 3; i < 20; i++ {
		a[i] = a[i-1] + a[i-2] + a[i-3]
	}

	fmt.Fprintf(writer, "%v\n", a[19])
}
