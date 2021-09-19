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
	var n int
	fmt.Fscan(reader, &n)
	s := make([][]string, 5)
	for i := 0; i < 5; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		s[i] = strings.Split(tmp, "")
	}

	ans := ""
	for i := 0; i < 4*n; i++ {
		if i%4 == 0 {
			continue
		}
		if i%4 == 1 {
			m := map[string]int{}
			for k := 0; k < 3; k++ {
				for j := 0; j < 5; j++ {
					m[s[j][i+k]]++
				}
			}
			switch m["#"] {
			case 7: // 7
				ans += "7"
			case 8: // 1
				ans += "1"
			case 9: // 4
				ans += "4"
			case 11: // 2, 3, 5
				switch {
				case s[1][i] == "#":
					ans += "5"
				case s[3][i] == "#":
					ans += "2"
				default:
					ans += "3"
				}
			case 12: // 0,6,9
				switch {
				case s[1][i+2] == ".":
					ans += "6"
				case s[3][i] == ".":
					ans += "9"
				default:
					ans += "0"
				}
			case 13: // 8
				ans += "8"
			}
		}

	}
	fmt.Fprintf(writer, "%v\n", ans)

}
