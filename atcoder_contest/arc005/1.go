package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		if i == n-1 {
			s[i] = strings.Split(s[i], ".")[0]
		}
	}

	m := map[string]int{}
	for i := 0; i < n; i++ {
		m[s[i]]++
	}
	fmt.Fprintf(writer, "%v\n", m["TAKAHASHIKUN"]+m["Takahashikun"]+m["takahashikun"])
}
