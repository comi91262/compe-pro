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
		if i%2 == 0 {
			ans += string(s[i])
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}

