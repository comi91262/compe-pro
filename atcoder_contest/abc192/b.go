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
		if i%2 == 0 && "a" <= s[i] && s[i] <= "z" {
			continue
		}
		if i%2 == 1 && "A" <= s[i] && s[i] <= "Z" {
			continue
		}
		fmt.Fprintf(writer, "%v\n", "No")
		return
	}
	fmt.Fprintf(writer, "%v\n", "Yes")
}
