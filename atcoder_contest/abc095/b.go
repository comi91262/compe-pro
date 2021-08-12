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
	var x int
	fmt.Fscan(reader, &x)

	var m = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &m[i])
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += m[i]
	}
	sort.Ints(m)

	fmt.Fprintf(writer, "%v\n", n+(x-sum)/m[0])

}
