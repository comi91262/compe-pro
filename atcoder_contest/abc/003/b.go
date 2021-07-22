package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func contains(x byte, a []byte) bool {
	for _, e := range a {
		if x == e {
			return true
		}
	}
	return false
}

func main() {
	defer writer.Flush()

	var s, t string
	fmt.Fscan(reader, &s, &t)

	for i := range s {
		if t[i] == '@' && s[i] == '@' {
			continue
		}

		if s[i] == '@' {
			if contains(t[i], []byte{byte('a'), byte('t'), byte('c'), byte('o'), byte('d'), byte('e'), byte('r')}) {
				continue
			}
		}

		if t[i] == '@' {
			if contains(s[i], []byte{byte('a'), byte('t'), byte('c'), byte('o'), byte('d'), byte('e'), byte('r')}) {
				continue
			}
		}

		if s[i] != t[i] {
			fmt.Fprintf(writer, "%v\n", "You will lose")
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", "You can win")
}
