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

	line, _, _ := reader.ReadLine()
	ss := strings.Split(string(line), " ")

	m := map[string]int{}
	for _, s := range ss {
		on := false
		c := ""
		for i := 0; i < len(s); i++ {
			if s[i] == "@"[0] {
				on = true
				m[c]++
				c = ""
				continue
			}
			if on && s[i] != "@"[0] {
				c += string(s[i])
			}
		}
		m[c]++
	}

	ans := []string{}
	for k := range m {
		if k == "" {
			continue
		}
		ans = append(ans, k)
	}

	sort.Strings(ans)
	for i := 0; i < len(ans); i++ {
		fmt.Fprintf(writer, "%v\n", ans[i])
	}

}
