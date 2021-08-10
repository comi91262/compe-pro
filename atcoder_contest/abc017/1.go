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

	var s1 int
	fmt.Fscan(reader, &s1)
	var e1 int
	fmt.Fscan(reader, &e1)
	var s2 int
	fmt.Fscan(reader, &s2)
	var e2 int
	fmt.Fscan(reader, &e2)
	var s3 int
	fmt.Fscan(reader, &s3)
	var e3 int
	fmt.Fscan(reader, &e3)

	fmt.Fprintf(writer, "%v\n", s1/10*e1+s2/10*e2+s3/10*e3)

}
