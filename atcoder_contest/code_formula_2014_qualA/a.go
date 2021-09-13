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

	for i := 0; i*i*i <= a; i++ {
		if i*i*i == a {
			fmt.Fprintf(writer, "%v\n", "YES")
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", "NO")
}
