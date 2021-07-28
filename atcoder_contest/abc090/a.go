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
	var t string
	fmt.Fscan(reader, &t)
	var u string
	fmt.Fscan(reader, &u)
	fmt.Fprintf(writer, "%v\n", string(s[0])+string(t[1])+string(u[2]))
}
