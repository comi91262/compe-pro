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

type pair struct {
	first  string
	second string
}

func countZero(s string) (cnt int) {
	for i := range s {
		if s[i] == "0"[0] {
			cnt++
		} else {
			return
		}
	}
	return
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var s = make([]pair, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &s[i].first)
		s[i].second = strings.TrimLeft(s[i].first, "0")
	}

	sort.Slice(s, func(i, j int) bool {
		if len(s[i].second) == len(s[j].second) {
			if s[i].second == s[j].second {
				return countZero(s[i].first) > countZero(s[j].first)
			}
			return s[i].second < s[j].second
		} else {
			return len(s[i].second) < len(s[j].second)
		}
	})

	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%v\n", s[i].first)
	}
}
