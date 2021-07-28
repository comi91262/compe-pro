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
		switch s[i] - 48 {
		case 1:
			ans += string(s[i] + 8)
		case 9:
			ans += string(s[i] - 8)
		default:
			ans += string(s[i])
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
