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

	var s string
	fmt.Fscan(reader, &s)

	x1 := s[0] - 48
	x2 := s[1] - 48
	x3 := s[2] - 48
	x4 := s[3] - 48

	if x1 == x2 && x2 == x3 && x3 == x4 {
		fmt.Fprintf(writer, "%v\n", "Weak")
		return
	}

	if (x2 == x1+1 || x2 == 0 && x1 == 9) && (x3 == x2+1 || x3 == 0 && x2 == 9) && (x4 == x3+1 || x4 == 0 && x3 == 9) {
		fmt.Fprintf(writer, "%v\n", "Weak")
		return
	}

	fmt.Fprintf(writer, "%v\n", "Strong")
}
