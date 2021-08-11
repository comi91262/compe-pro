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

	mx := a[0]
	a = a[1:]
	sum := 0
	for i := range a {
		sum += a[i]
	}
	if mx < sum {
        fmt.Fprintf(writer, "%v\n", "Yes")
	} else {
        fmt.Fprintf(writer, "%v\n", "No")
    }
}
