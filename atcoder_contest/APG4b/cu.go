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
	fmt.Fprintf(writer, "%v\n%v\n", "いつも2525", "AtCoderくん")
}
