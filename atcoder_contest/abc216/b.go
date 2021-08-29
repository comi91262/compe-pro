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
	var t = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
		fmt.Fscan(reader, &t[i])
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if s[i] == s[j] && t[i] == t[j] {
				fmt.Fprintf(writer, "%v\n", "Yes")
				return
			}
		}
	}

    fmt.Fprintf(writer, "%v\n", "No")
}
