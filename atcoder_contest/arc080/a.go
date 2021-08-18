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
		switch {
		case a[i]%4 == 0:
			m[4]++
		case a[i]%2 == 0:
			m[2]++
		default:
			m[1]++
		}
	}

	if m[4]+1 >= m[2]%2+m[1] {
		fmt.Fprintf(writer, "%v\n", "Yes")
	} else {

		fmt.Fprintf(writer, "%v\n", "No")
	}
}
