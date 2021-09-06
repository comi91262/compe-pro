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

	a := []int{}
	for i := 0; i < 8; i++ {
		var t int
		fmt.Fscan(reader, &t)
		a = append(a, t)
	}
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
	fmt.Fprintf(writer, "%v\n", a[0])

}
