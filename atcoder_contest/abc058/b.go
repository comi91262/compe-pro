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
	var o string
	fmt.Fscan(reader, &o)
	var e string
	fmt.Fscan(reader, &e)

	ans := ""
	for i := 0; i < len(o); i++ {
		ans += string(o[i])
		if i != len(e) {
			ans += string(e[i])
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)

}
