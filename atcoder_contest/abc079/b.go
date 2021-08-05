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

	var a = make([]int, n+1)

	a[0] = 2
	a[1] = 1
	for i := 2; i <= n; i++ {
		a[i] = a[i-1] + a[i-2]
	}
	fmt.Fprintf(writer, "%v\n", a[n])

}
