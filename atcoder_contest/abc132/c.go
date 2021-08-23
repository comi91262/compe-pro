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
	var d = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &d[i])
	}
	sort.Ints(d)
	fmt.Fprintf(writer, "%v\n", d[len(d)/2]-d[len(d)/2-1])

}
