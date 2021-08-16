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

	var s string
	fmt.Fscan(reader, &s)
	var k int
	fmt.Fscan(reader, &k)

	if len(s) < k {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}

	m := map[string]int{}
	for i := 0; i < len(s)-k+1; i++ {
		m[s[i:i+k]]++
	}
	fmt.Fprintf(writer, "%v\n", len(m))
}
