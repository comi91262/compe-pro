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

	if 1 <= a && a <= 9 && 1 <= b && b <= 9 {
		fmt.Fprintf(writer, "%v\n", a*b)
	} else {
		fmt.Fprintf(writer, "%v\n", -1)
	}

}
