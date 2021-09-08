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
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	n := len(s)
	for i := 0; i < n; i++ {
		if s[i] != s[n-i-1] {
			fmt.Fprintf(writer, "%v\n", "NO")
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", "YES")
}
