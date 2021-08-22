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

	if n%2 == 1 {
		fmt.Fprintf(writer, "%v\n", "No")
		return
	}

	for i := 0; i < n/2; i++ {
		if s[i] != s[n/2+i] {
			fmt.Fprintf(writer, "%v\n", "No")
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", "Yes")
}
