package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	var s1, s2, s3, s4 string
	fmt.Fscan(reader, &s1, &s2, &s3, &s4)
	s := []string{s1, s2, s3, s4}
	t := []string{"H", "2B", "3B", "HR"}

	sort.Strings(s)
	sort.Strings(t)

	if eq(s, t) {
		fmt.Fprintf(writer, "%v\n", "Yes")
	} else {
		fmt.Fprintf(writer, "%v\n", "No")
	}
}
