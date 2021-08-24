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

	var x string
	fmt.Fscan(reader, &x)
	fmt.Fprintf(writer, "%v\n", strings.Split(x, ".")[0])
}

