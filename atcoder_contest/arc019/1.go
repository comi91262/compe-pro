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

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case "O":
			fmt.Fprintf(writer, "0")
		case "D":
			fmt.Fprintf(writer, "0")
		case "I":
			fmt.Fprintf(writer, "1")
		case "Z":
			fmt.Fprintf(writer, "2")
		case "S":
			fmt.Fprintf(writer, "5")
		case "B":
			fmt.Fprintf(writer, "8")
		default:
			fmt.Fprintf(writer, "%v", s[i])
		}
	}
	fmt.Fprintf(writer, "\n")
}
