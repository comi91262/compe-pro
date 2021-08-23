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

	var n int
	fmt.Fscan(reader, &n)
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	ans := 0
	for i := 0; i < len(s)-2; i++ {
		if s[i] == "A" && s[i+1] == "B" && s[i+2] == "C" {
			ans++
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
