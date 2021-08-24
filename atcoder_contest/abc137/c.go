package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	var s = make([][]string, n)
	for i := 0; i < n; i++ {
		var ss string
		fmt.Fscan(reader, &ss)
		s[i] = strings.Split(ss, "")
	}
	m := map[string]int{}
	for i := 0; i < n; i++ {
		sort.Slice(s[i], func(k, j int) bool { return s[i][k] < s[i][j] })
		m[strings.Join(s[i], "")]++
	}
	ans := 0
	for _, v := range m {
		ans += v * (v - 1) / 2
	}

	fmt.Fprintf(writer, "%v\n", ans)

}
