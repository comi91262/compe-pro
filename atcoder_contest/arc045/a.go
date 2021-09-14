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

	var s []string
	for i := 0; i < 100; i++ {
		var ss string
		fmt.Fscan(reader, &ss)
		if ss == "" {
			continue
		}
		s = append(s, ss)
	}

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case "Left":
			fmt.Fprintf(writer, "%v", "<")
		case "Right":
			fmt.Fprintf(writer, "%v", ">")
		case "AtCoder":
			fmt.Fprintf(writer, "%v", "A")
		}
		if i != len(s)-1 {
			fmt.Fprintf(writer, " ")
		}
	}
	fmt.Fprintf(writer, "\n")

}
