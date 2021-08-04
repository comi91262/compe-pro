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

	fmt.Fprintf(writer, "%v %v %v %v %v %v\n", m["A"], m["B"], m["C"], m["D"], m["E"], m["F"])
}
