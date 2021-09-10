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
	var a = make([]int, 6)
	for i := 0; i < 6; i++ {
		fmt.Fscan(reader, &a[i])
	}
	sort.Slice(a, func(i, j int) bool { return a[i] > a[j] })
	fmt.Fprintf(writer, "%v\n", a[3-1])
}
