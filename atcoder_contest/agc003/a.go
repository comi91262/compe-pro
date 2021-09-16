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
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	if m["N"] > 0 && m["S"] == 0 {
		fmt.Fprintf(writer, "%v\n", "No")
		return
	}
	if m["S"] > 0 && m["N"] == 0 {
		fmt.Fprintf(writer, "%v\n", "No")
		return
	}
	if m["E"] > 0 && m["W"] == 0 {
		fmt.Fprintf(writer, "%v\n", "No")
		return
	}
	if m["W"] > 0 && m["E"] == 0 {
		fmt.Fprintf(writer, "%v\n", "No")
		return
	}
	fmt.Fprintf(writer, "%v\n", "Yes")
}
