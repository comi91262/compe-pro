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

	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })

	s, t := 0, 0
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			s += a[i]
		} else {
			t += a[i]
		}
	}

	fmt.Fprintf(writer, "%v\n", s)
}
