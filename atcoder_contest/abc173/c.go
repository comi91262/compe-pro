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
	var k int
	fmt.Fscan(reader, &k)

	s := make([][]string, h)
	for i := 0; i < h; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		s[i] = strings.Split(tmp, "")
	}

	t := make([][]string, h)
	for i := 0; i < h; i++ {
		t[i] = make([]string, w)
	}
	ans := 0
	for i := 0; i < 1<<h; i++ {
		for j := 0; j < 1<<w; j++ {
			for u := 0; u < h; u++ {
				for v := 0; v < w; v++ {
					t[u][v] = s[u][v]
				}
			}

			for u := 0; u < h; u++ {
				if i&(1<<u) != 0 {
					for v := 0; v < w; v++ {
						t[u][v] = "x"
					}
				}
			}
			for u := 0; u < w; u++ {
				if j&(1<<u) != 0 {
					for v := 0; v < h; v++ {
						t[v][u] = "x"
					}
				}
			}

			cnt := 0
			for u := 0; u < h; u++ {
				for v := 0; v < w; v++ {
					if t[u][v] == "#" {
						cnt++
					}
				}
			}

			if k == cnt {
				ans++
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", ans)
}
