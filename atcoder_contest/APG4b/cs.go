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
	fmt.Fprintf(writer, "%v\n", 365*24*60*60*1)
	fmt.Fprintf(writer, "%v\n", 365*24*60*60*2)
	fmt.Fprintf(writer, "%v\n", 365*24*60*60*5)
	fmt.Fprintf(writer, "%v\n", 365*24*60*60*10)
}
