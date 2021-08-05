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

	m := map[string]int{}
	for i := range s {
		m[s[i]]++
	}
	if m["A"] == 3 || m["B"] == 3 {
		fmt.Fprintf(writer, "%v\n", "No")
	} else {
		fmt.Fprintf(writer, "%v\n", "Yes")
	}
}
