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

	n := 12
	k := 4
	cnt := 0
	for i := 0; i < n-k+1; i++ {
		if s[i] == "Z" && s[i+1] == "O" && s[i+2] == "N" && s[i+3] == "e" {
			cnt++
		}
	}
	fmt.Fprintf(writer, "%v\n", cnt)
}
