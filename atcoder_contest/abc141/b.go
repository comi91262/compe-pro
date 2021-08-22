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

	for i := range s {
		if i%2 == 0 {
			if s[i] == "R" || s[i] == "U" || s[i] == "D" {
				continue
			}
		} else {
			if s[i] == "L" || s[i] == "U" || s[i] == "D" {
				continue
			}
		}

		fmt.Fprintf(writer, "%v\n", "No")
		return
	}
	fmt.Fprintf(writer, "%v\n", "Yes")
}
