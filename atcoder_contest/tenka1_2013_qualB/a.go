package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var a = make([]string, 15)
	for i := 0; i < 15; i++ {
		fmt.Fscan(reader, &a[i])
	}
	sort.Strings(a)
	fmt.Fprintf(writer, "%v\n", a[7-1])
}
