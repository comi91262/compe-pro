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

	var h int
	fmt.Fscan(reader, &h)
	var w int
	fmt.Fscan(reader, &w)

	var s = make([][]string, h)
	for i := 0; i < h; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		s[i] = strings.Split(tmp, "")
	}

	ans := 0
	for i := 0; i < h-2+1; i++ {
		for j := 0; j < w-2+1; j++ {
			m := map[string]int{}
			m[s[i][j]]++
			m[s[i][j+1]]++
			m[s[i+1][j]]++
			m[s[i+1][j+1]]++

			if m["#"] == 3 && m["."] == 1 || m["#"] == 1 && m["."] == 3 {
				ans++
			}
		}
	}

	fmt.Fprintf(writer, "%v\n", ans)
}
