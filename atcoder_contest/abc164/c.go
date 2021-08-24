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

	m := map[string]int{}
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		m[tmp]++
	}
	fmt.Fprintf(writer, "%v\n", len(m))
}
