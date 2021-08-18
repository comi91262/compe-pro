package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func sum(a []int) int {
	r := 0
	for i := range a {
		r += a[i]
	}
	return r
}
func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	total := sum(a)
	if total%10 != 0 {
		fmt.Fprintf(writer, "%v\n", total)
		return
	}

	sort.Ints(a)
	for i := 0; i < n; i++ {
		if a[i]%10 != 0 {
			fmt.Fprintf(writer, "%v\n", total-a[i])
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", 0)
}
