package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	ans := 1
	for i := 0; i < n-1; i++ {
		if s[i] != s[i+1] {
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
