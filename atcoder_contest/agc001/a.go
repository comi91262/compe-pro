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
	var l = make([]int, 2*n)
	for i := 0; i < 2*n; i++ {
		fmt.Fscan(reader, &l[i])
	}
	sort.Ints(l)

	ans := 0
	for i := 0; i < len(l); i += 2 {
		ans += l[i]
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
