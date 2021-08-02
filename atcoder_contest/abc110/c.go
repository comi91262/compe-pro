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

	var ss string
	fmt.Fscan(reader, &ss)
	var tt string
	fmt.Fscan(reader, &tt)

	s := strings.Split(ss, "")
	t := strings.Split(tt, "")

	m := map[string]string{}
	for i := range s {
		if m[s[i]] == "" {
			m[s[i]] = t[i]
		} else if m[s[i]] != t[i] {
			fmt.Fprintf(writer, "%v\n", "No")
			return
		}
	}

	s, t = t, s

	m = map[string]string{}
	for i := range t {
		if m[s[i]] == "" {
			m[s[i]] = t[i]
		} else if m[s[i]] != t[i] {
			fmt.Fprintf(writer, "%v\n", "No")
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", "Yes")
}
