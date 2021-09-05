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

	if len(s) == 2 && s[0] == s[1] {
		fmt.Fprintf(writer, "%v %v\n", 1, 2)
		return
	}

	for i := 0; i < len(s)-3+1; i++ {
		a, b, c := s[i], s[i+1], s[i+2]
		if a == c {
			fmt.Fprintf(writer, "%v %v\n", i+1, i+3)
			return
		}
		if a == b {
			fmt.Fprintf(writer, "%v %v\n", i+1, i+2)
			return
		}
		if b == c {
			fmt.Fprintf(writer, "%v %v\n", i+2, i+3)
			return
		}
	}
	fmt.Fprintf(writer, "%v %v\n", -1, -1)
}
