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

	if len(s)%2 != 0 {
		fmt.Fprintf(writer, "%v\n", "No")
		return
	}

	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			if s[i] != "h" {
				fmt.Fprintf(writer, "%v\n", "No")
				return
			}
		} else {
			if s[i] != "i" {
				fmt.Fprintf(writer, "%v\n", "No")
				return
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", "Yes")
}
