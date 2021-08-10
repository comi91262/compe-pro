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

	var s = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
	}

	m := map[string]int{}
	for i := 0; i < n; i++ {
		m[s[i]]++
	}

	ans := ""
	mx := 0
	for s, v := range m {
		if v > mx {
			mx = v
			ans = s
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
