package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func eq(a, b []string) bool {

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}

	}
	return true
}

func main() {
	defer writer.Flush()

	var ss, tt string
	fmt.Fscan(reader, &ss)
	fmt.Fscan(reader, &tt)

	if ss == tt {
		fmt.Fprintf(writer, "%v\n", "Yes")
		return
	}

	var s = strings.Split(ss, "")
	var t = strings.Split(tt, "")

	for i := 0; i < len(s)-1; i++ {
		u := make([]string, len(s))
		copy(u, s)
		u[i], u[i+1] = u[i+1], u[i]
		if eq(t, u) {
			fmt.Fprintf(writer, "%v\n", "Yes")
			return
		}
	}

	fmt.Fprintf(writer, "%v\n", "No")
}
