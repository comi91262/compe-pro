package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)
	var w = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &w[i])
	}

	var pre string
	m := map[string]int{}
	for i := 0; i < n; i++ {
		if i == 0 {
			m[w[i]]++
			pre = w[i]
			continue
		}

		if m[w[i]] > 0 || pre[len(pre)-1] != w[i][0] {
			fmt.Fprintf(writer, "%v\n", "No")
			return
		}
		m[w[i]]++
		pre = w[i]
	}

	fmt.Fprintf(writer, "%v\n", "Yes")
}
