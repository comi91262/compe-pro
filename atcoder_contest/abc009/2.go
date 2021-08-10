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
	m := map[int]int{}
	for i := 0; i < n; i++ {
		m[a[i]]++
	}

	var b = make([]int, 0, len(m))
	for k := range m {
		b = append(b, k)
	}
	sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })

	fmt.Fprintf(writer, "%v\n", b[1])

}
