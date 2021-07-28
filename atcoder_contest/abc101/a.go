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
	for i := range s {
		switch s[i] {
		case "+"[0]:
			ans++
		case "-"[0]:
			ans--
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
