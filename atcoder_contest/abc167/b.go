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
	var b int
	fmt.Fscan(reader, &b)
	var c int
	fmt.Fscan(reader, &c)
	var k int
	fmt.Fscan(reader, &k)

	if k <= a {
		fmt.Fprintf(writer, "%v\n", k)
		return
	}
	if k <= a+b {
		fmt.Fprintf(writer, "%v\n", a)
		return
	}

	if k <= a+b+c {
		fmt.Fprintf(writer, "%v\n", a-(k-a-b))
		return
	}

}
