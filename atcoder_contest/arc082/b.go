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
	var p = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &p[i])
	}

	// fmt.Fprintf(writer, "%v\n", p)
	for i := 0; i < n; i++ {
		if p[i] != i+1 {
			p[i] = 0
		} else {
			p[i] = 1
		}
	}

	// fmt.Fprintf(writer, "%v\n", p)

	ans := 0
	for i := 0; i < n; i++ {
		if p[i] == 1 {
			if i-1 >= 0 && p[i-1] == 1 {
				p[i] = 0
				p[i-1] = 0
				ans++
				continue
			}

			if i+1 <= n-1 && p[i+1] == 1 {
				p[i] = 0
				p[i+1] = 0
				ans++
				continue
			}
		}
	}

	for i := 0; i < n; i++ {
		if p[i] == 1 {
			ans++
		}
	}
	// fmt.Fprintf(writer, "%v\n", p)
	fmt.Fprintf(writer, "%v\n", ans)
}
