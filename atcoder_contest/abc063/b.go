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

	if len(s) == len(m) {
		fmt.Fprintf(writer, "%v\n", "yes")
	} else {
		fmt.Fprintf(writer, "%v\n", "no")
	}

}
