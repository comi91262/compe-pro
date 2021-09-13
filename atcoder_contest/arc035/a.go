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

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] && s[i] != "*" && s[len(s)-i-1] != "*" {
			fmt.Fprintf(writer, "%v\n", "NO")
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", "YES")
}
