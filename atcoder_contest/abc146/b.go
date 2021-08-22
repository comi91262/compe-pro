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

	var n int
	fmt.Fscan(reader, &n)

	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	for i := range s {
		fmt.Fprintf(writer, "%v", string("A"[0]+(s[i][0]+byte(n)-"A"[0])%26))
	}
	fmt.Fprintf(writer, "\n")

}
