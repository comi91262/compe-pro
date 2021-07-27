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

	var w string
	fmt.Fscan(reader, &w)

	m := map[byte]int{}
	for i := range w {
		m[w[i]]++
	}

	for i := range m {
		if m[i]%2 != 0 {
			fmt.Fprintf(writer, "%v\n", "No")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "Yes")
}
