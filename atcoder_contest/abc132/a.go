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

	var s string
	fmt.Fscan(reader, &s)

	m := map[byte]int{}

	for i := range s {
		m[s[i]]++
	}

	if len(m) != 2 {
		fmt.Fprintf(writer, "%v\n", "No")
		return
	}

	for _, v := range m {
		if v != 2 {
			fmt.Fprintf(writer, "%v\n", "No")
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", "Yes")
}
