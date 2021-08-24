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
	var d int
	fmt.Fscan(reader, &d)

	ans := 0
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(reader, &x)
		var y int
		fmt.Fscan(reader, &y)

		if x*x+y*y <= d*d {
			ans++
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
