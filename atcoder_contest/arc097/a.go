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

	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")
	var k int
	fmt.Fscan(reader, &k)

	ans := []string{}
	m := map[string]int{}
	for i := 1; i <= 5; i++ {
		for j := 0; j < len(s)-i+1; j++ {
			ss = strings.Join(s[j:j+i], "")
			if m[ss] == 0 {
				ans = append(ans, ss)
			}
			m[ss]++
		}
	}
	sort.Strings(ans)
	fmt.Fprintf(writer, "%v\n", ans[k-1])
}
