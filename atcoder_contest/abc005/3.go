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

	var t, n int
	fmt.Fscan(reader, &t, &n)

	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}

	var m int
	fmt.Fscan(reader, &m)

	var b = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscan(reader, &b[i])
	}

	if n < m {
		fmt.Fprintf(writer, "%v\n", "no")
		return
	}

	mapper := map[int]int{}

	for i := 0; i < m; i++ {
		for j := i; j < n; j++ {
			if a[j] <= b[i] && b[i] <= a[j]+t {
				if mapper[j] == 0 {
					mapper[j]++
					break
				}
			}
		}
	}

	if len(mapper) == m {
		fmt.Fprintf(writer, "%v\n", "yes")
	} else {
		fmt.Fprintf(writer, "%v\n", "no")
	}

}
