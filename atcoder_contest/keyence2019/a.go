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

	var x int
	m := map[int]int{}

	fmt.Fscan(reader, &x)
	m[x]++
	fmt.Fscan(reader, &x)
	m[x]++
	fmt.Fscan(reader, &x)
	m[x]++
	fmt.Fscan(reader, &x)
	m[x]++

	if m[1]*m[9]*m[7]*m[4] == 1 {
      fmt.Fprintf(writer, "%v\n", "YES")
	} else {
		fmt.Fprintf(writer, "%v\n", "NO")
	}
}
