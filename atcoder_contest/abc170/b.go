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
	var x int
	fmt.Fscan(reader, &x)
	var y int
	fmt.Fscan(reader, &y)

	for i := 0; i <= x; i++ {
		if i*2+(x-i)*4 == y {
			fmt.Fprintf(writer, "%v\n", "Yes")
			return
		}

	}
	fmt.Fprintf(writer, "%v\n", "No")
}
