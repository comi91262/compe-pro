package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func check(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)/2+i] {
			return false
		}
	}

	return true
}

func main() {
	defer writer.Flush()

	var s string
	fmt.Fscan(reader, &s)

	s = s[:len(s)-1]
	for {
		if check(s) {
			fmt.Fprintf(writer, "%v\n", len(s))
			return
		}

		s = s[:len(s)-1]
		if len(s) == 1 {
			fmt.Fprintf(writer, "%v\n", 1)
			return
		}
	}

}

