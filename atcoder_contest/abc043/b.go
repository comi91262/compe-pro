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
		case "0"[0]:
			ans += "0"
		case "1"[0]:
			ans += "1"
		case "B"[0]:
			if len(ans) == 0 {
				continue
			}
			ans = ans[0 : len(ans)-1]
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
