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

	var a, b, c int
	fmt.Fscan(reader, &a, &b, &c)

	m := map[int]int{}

	m[a]++
	m[b]++
	m[c]++

	if m[5] == 2 && m[7] == 1 {
		fmt.Fprintf(writer, "%v\n", "YES")
	} else {
		fmt.Fprintf(writer, "%v\n", "NO")
	}

}
