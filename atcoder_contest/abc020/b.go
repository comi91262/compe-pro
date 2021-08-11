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

	var a string
	fmt.Fscan(reader, &a)
	var b string
	fmt.Fscan(reader, &b)
	i, _ := strconv.Atoi(a + b)
	fmt.Fprintf(writer, "%v\n", i*2)
}
