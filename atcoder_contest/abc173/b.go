package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var n int
	fmt.Fscan(reader, &n)

	m := map[string]int{}
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(reader, &s)
		m[s]++
	}
	fmt.Fprintf(writer, "AC x %v\n", m["AC"])
	fmt.Fprintf(writer, "WA x %v\n", m["WA"])
	fmt.Fprintf(writer, "TLE x %v\n", m["TLE"])
	fmt.Fprintf(writer, "RE x %v\n", m["RE"])
}