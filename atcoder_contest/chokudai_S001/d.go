package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	sort.Ints(a)
	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%v", a[i])
		if i != n-1 {
			fmt.Fprintf(writer, " ")
		}
	}
	fmt.Fprintf(writer, "\n")
}
