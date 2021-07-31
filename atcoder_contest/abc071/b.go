package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var s string
	fmt.Fscan(reader, &s)

	t := strings.Split(s, "")
	sort.Strings(t)

	m := map[string]int{}
	for i := range t {
		m[t[i]]++
	}
	if len(m) == 26 {
		fmt.Fprintf(writer, "%v\n", "None")
		return
	}

	s2 := make([]string, len(m))
	index := 0
	for key := range m {
		s2[index] = key
		index++
	}
	sort.Strings(s2)

	s3 := "abcdefghijklmnopqrstuvwxyz"
	t3 := strings.Split(s3, "")

	idx := 0
	for i := 0; i < len(s2); i++ {
		if s2[i] != t3[i] {
			fmt.Fprintf(writer, "%v\n", t3[i])
			return
		}
		idx++
	}
	fmt.Fprintf(writer, "%v\n", t3[idx])
}
