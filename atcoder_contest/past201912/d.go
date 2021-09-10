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
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &a[i])
	}
	m := map[int]int{}
	for i := 0; i < n; i++ {
		m[a[i]]++
	}

	if len(m) == n {
		fmt.Fprintf(writer, "%v\n", "Correct")
		return
	}

	l := -1
	for i := 0; i < n; i++ {
		if m[i+1] == 0 {
			l = i + 1
			break
		}
	}

	for i := 0; i < n; i++ {
		if m[i+1] > 1 {
			fmt.Fprintf(writer, "%v %v\n", i+1, l)
			break
		}
	}

}
