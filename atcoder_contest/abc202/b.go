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
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}

	for i := range s {
		switch s[i] {
		case "6":
			s[i] = "9"
		case "8":
			s[i] = "8"
		case "9":
			s[i] = "6"
		}
	}
	fmt.Fprintf(writer, "%v\n", strings.Join(s, ""))
}
