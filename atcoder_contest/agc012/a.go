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

	var a = make([]int, 3*n)
	for i := 0; i < 3*n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	sort.Ints(a)
	ans := 0
	for i := n; i < 3*n; i++ {
		if (i-n)%2 == 0 {
			ans += a[i]
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
