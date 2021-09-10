package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var s string
	fmt.Fscan(reader, &s)

	if n, err := strconv.Atoi(s); err == nil {
		fmt.Fprintf(writer, "%v\n", n*2)
	} else {
		fmt.Fprintf(writer, "%v\n", "error")
	}
}
