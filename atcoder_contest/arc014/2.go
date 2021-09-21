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
	var a = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	m := map[string]int{}
	pre := a[0]
	m[a[0]]++
	for i := 1; i < n; i++ {
		first := i%2 == 0

		if m[a[i]] > 0 {
			if first {
				fmt.Fprintf(writer, "%v\n", "LOSE")
			} else {
				fmt.Fprintf(writer, "%v\n", "WIN")
			}
			return
		}
		m[a[i]]++

		if pre[len(pre)-1] != a[i][0] {
			if first {
				fmt.Fprintf(writer, "%v\n", "LOSE")
			} else {
				fmt.Fprintf(writer, "%v\n", "WIN")
			}
			return
		}
		pre = a[i]
	}
	fmt.Fprintf(writer, "%v\n", "DRAW")
}
