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

	for i := 0; i < n-1; i++ {
		if a[i+1] >= a[i] {
			continue
		}

		if a[i+1] >= a[i]-1 {
			a[i+1]++
			continue
		}

		fmt.Fprintf(writer, "%v\n", "No")
		return
	}
	fmt.Fprintf(writer, "%v\n", "Yes")
}
