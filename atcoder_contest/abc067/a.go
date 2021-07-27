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

	if a%3 == 0 || b%3 == 0 || (a+b)%3 == 0 {
		fmt.Fprintf(writer, "%v\n", "Possible")
	} else {
		fmt.Fprintf(writer, "%v\n", "Impossible")
	}

}
