package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var s string
	fmt.Fscan(reader, &s)

	ss := strings.Split(s, "/")

	if ss[1] <= "03" || ss[1] == "04" && ss[2] <= "30" {
		fmt.Fprintf(writer, "%v\n", "Heisei")
	} else {
		fmt.Fprintf(writer, "%v\n", "TBD")

	}
}

