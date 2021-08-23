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

	for i := 0; i < n; i++ {
		if a[i]%2 == 0 {
			if a[i]%3 != 0 && a[i]%5 != 0 {
				fmt.Fprintf(writer, "%v\n", "DENIED")
				return
			}
		}

	}
	fmt.Fprintf(writer, "%v\n", "APPROVED")
}
