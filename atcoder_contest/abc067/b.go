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
	var k int
	fmt.Fscan(reader, &k)

	var l = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &l[i])
	}

	sort.Ints(l)

	sum := 0
	for i := 0; i < k; i++ {
		sum += l[len(l)-1-i]
	}
	fmt.Fprintf(writer, "%v\n", sum)

}
