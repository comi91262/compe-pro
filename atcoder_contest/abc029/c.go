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

	bitfullSearch(3, n)
}

func bitfullSearch(base int, n int) {
	power := make([]int, n+1)
	power[0] = 1
	for i := 0; i < n; i++ {
		power[i+1] = power[i] * base
	}

	for i := 0; i < power[n]; i++ {
		for j := 0; j < n; j++ {
			bit := i / power[n-1-j] % base
			switch bit {
			case 0:
				fmt.Fprintf(writer, "%v", "a")
			case 1:
				fmt.Fprintf(writer, "%v", "b")
			case 2:
				fmt.Fprintf(writer, "%v", "c")
			}
		}
		fmt.Fprintf(writer, "\n")
	}
}
