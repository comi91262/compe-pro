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
	var a int
	fmt.Fscan(reader, &a)
	var op string
	fmt.Fscan(reader, &op)
	var c int
	fmt.Fscan(reader, &c)

	if op == "+" {
		fmt.Fprintf(writer, "%v\n", a+c)

	} else {
		fmt.Fprintf(writer, "%v\n", a-c)
	}

}
