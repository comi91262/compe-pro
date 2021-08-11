package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func f(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return 3*n + 1
	}

}
func main() {
	defer writer.Flush()
	var s int
	fmt.Fscan(reader, &s)

	m := map[int]int{}
	idx := 1
	m[s]++
	for {
		s = f(s)
		idx++
		if m[s] > 0 {
			fmt.Fprintf(writer, "%v\n", idx)
			return
		}
		m[s]++
	}

}
