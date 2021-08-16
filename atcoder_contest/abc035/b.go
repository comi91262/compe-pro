package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}
func main() {
	defer writer.Flush()

	var ss string
	fmt.Fscan(reader, &ss)
	var s = strings.Split(ss, "")
	var t int
	fmt.Fscan(reader, &t)

	m := map[string]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	ans := 0
	switch t {
	case 1:
		ans = abs(m["U"]-m["D"]) + abs(m["L"]-m["R"]) + m["?"]
	case 2:
		total := abs(m["U"]-m["D"]) + abs(m["L"]-m["R"])
		if total >= m["?"] {
			ans = total - m["?"]
		} else {
			total = m["?"] - total
			switch total % 4 {
			case 0:
				ans = 0
			case 1:
				ans = 1
			case 2:
				ans = 0
			case 3:
				ans = 1
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)

}
