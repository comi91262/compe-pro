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

	ans := ""
	for i := range s {
		switch s[i] {
		case 97:
		case 105:
		case 111:
		case 101:
		case 117:
		default:
			ans += string(s[i])
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)

}
