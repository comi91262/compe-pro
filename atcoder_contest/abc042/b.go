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
	var l int
	fmt.Fscan(reader, &l)

	var s = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
	}

	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

	ans := ""
	for i := 0; i < n; i++ {
		ans += s[i]
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
