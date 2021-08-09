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

	if a%1000 == 0 {
		fmt.Fprintf(writer, "%v\n", 0)
	} else {
		fmt.Fprintf(writer, "%v\n", (a+1000)/1000*1000-a)
	}
}


