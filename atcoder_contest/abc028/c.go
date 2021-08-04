package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

var r []int

func dfs(m, sum int, a []int) {
	if m == 0 || len(a) == 0 {
		r = append(r, sum)
		return
	}

	dfs(m-1, sum+a[0], a[1:])
	dfs(m, sum, a[1:])
}

func main() {
	defer writer.Flush()

	var a = make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Fscan(reader, &a[i])
	}
	dfs(3, 0, a)

	sort.Slice(r, func(i, j int) bool { return r[i] > r[j] })
	fmt.Fprintf(writer, "%v\n", r[2])
}
