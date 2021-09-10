package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func toLower(s string) string {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		r[i] = unicode.ToLower(r[i])
	}

	return string(r)
}

func main() {
	defer writer.Flush()
	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")

	ans := []string{}
	cnt := 0
	stack := ""
	for i := 0; i < len(s); i++ {
		if "A" <= s[i] && s[i] <= "Z" {
			cnt++
		}
		stack += s[i]
		if cnt == 2 {
			cnt = 0
			ans = append(ans, stack)
			stack = ""
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return toLower(ans[i]) < toLower(ans[j])
	})

	fmt.Fprintf(writer, "%v\n", strings.Join(ans, ""))
}
