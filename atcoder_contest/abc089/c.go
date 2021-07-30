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

	var n int
	fmt.Fscan(reader, &n)

	var s = make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i])
	}

	m := map[string]int{}

	for i := range s {
		switch s[i][0] {
		case "M"[0]:
			m["M"]++
		case "A"[0]:
			m["A"]++
		case "R"[0]:
			m["R"]++
		case "C"[0]:
			m["C"]++
		case "H"[0]:
			m["H"]++
		default:
		}
	}

	t := [][]string{
		{"M", "A", "R"},
		{"M", "A", "C"},
		{"M", "A", "H"},
		{"M", "R", "C"},
		{"M", "R", "H"},
		{"M", "C", "H"},
		{"A", "R", "C"},
		{"A", "R", "H"},
		{"A", "C", "H"},
		{"R", "C", "H"},
	}

	ans := 0
	for i := range t {
		cnt := 1
		for j := range t[i] {
			cnt *= m[t[i][j]]
		}
		ans += cnt
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
