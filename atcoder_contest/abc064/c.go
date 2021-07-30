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

	acn := 0
	for i := 0; i < n; i++ {
		if a[i] >= 3200 {
			acn++
			continue
		}
		m[a[i]/400]++
	}

	cn := len(m)

	if cn == 0 {
		fmt.Fprintf(writer, "%v %v\n", 1, acn)
	} else {
		fmt.Fprintf(writer, "%v %v\n", cn, cn+acn)
	}
}
