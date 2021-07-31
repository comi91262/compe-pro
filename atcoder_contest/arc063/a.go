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

	ans := 0
	for i := 0; i < len(s)-1; i++ {
		if s[i+1] != s[i] {
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
