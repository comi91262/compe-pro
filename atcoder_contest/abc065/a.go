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
	fmt.Fscan(reader, &x)
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

	if b-a <= 0 {
		fmt.Fprintf(writer, "%v\n", "delicious")
		return
	}
	if b-a <= x {
		fmt.Fprintf(writer, "%v\n", "safe")
	} else {
		fmt.Fprintf(writer, "%v\n", "dangerous")
	}

}
