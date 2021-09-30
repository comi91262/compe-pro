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

	if s[0] != s[n-1] {
		fmt.Fprintf(writer, "%v\n", 1)
		return
	}
	for i := 1; i < n-2; i++ {
		if s[0] != s[i] && s[i+1] != s[n-1] && s[i] == s[i+1] {
			fmt.Fprintf(writer, "%v\n", 2)
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", -1)
}
