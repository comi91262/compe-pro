package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type pair struct {
	x string
	y int
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var s = make([]string, n)
	var t = make([]int, n)
	var u = make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
		fmt.Fscan(reader, &t[i])
		u[i] = pair{s[i], t[i]}
	}

	sort.Slice(u, func(i, j int) bool { return u[i].y > u[j].y })
	fmt.Fprintf(writer, "%v\n", u[1].x)
}
