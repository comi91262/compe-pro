package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func testEqSlice(a, b []string) bool {

	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}

	}
	return true
}
func testEqSlice2(a, b [][]string) bool {

	n := len(a)
	m := len(b)
	if n != m {
		return false
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i != j {
				continue
			}
			if !testEqSlice(a[i], b[i]) {
				return false
			}
		}
	}

	return true
}

func main() {
	defer writer.Flush()
	var h int
	fmt.Fscan(reader, &h)
	var w int
	fmt.Fscan(reader, &w)

	s := make([][]string, h)
	for i := 0; i < h; i++ {
		s[i] = make([]string, w)
		var tmp string
		fmt.Fscan(reader, &tmp)
		s[i] = strings.Split(tmp, "")
	}
	var t = make([][]string, h)
	for i := 0; i < h; i++ {
		t[i] = make([]string, w)
		for j := 0; j < w; j++ {
			t[i][j] = "."

		}
	}

	var dx = [8]int{1, 0, -1, 0, 1, 1, -1, -1}
	var dy = [8]int{0, 1, 0, -1, 1, -1, 1, -1}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ok := true
			for k := 0; k < 8; k++ {
				x := i + dx[k]
				y := j + dy[k]

				if 0 <= x && x < h && 0 <= y && y < w {
					if s[x][y] != "#" {
						ok = false
						break

					}
				}
			}
			if ok && s[i][j] == "#" {
				t[i][j] = "#"
			}
		}
	}

	var u = make([][]string, h)
	for i := 0; i < h; i++ {
		u[i] = make([]string, w)
		for j := 0; j < w; j++ {
			u[i][j] = "."

		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ok := false
			for k := 0; k < 8; k++ {
				x := i + dx[k]
				y := j + dy[k]

				if 0 <= x && x < h && 0 <= y && y < w {
					if t[x][y] == "#" {
						ok = true
						break
					}
				}
			}
			if ok || t[i][j] == "#" {
				u[i][j] = "#"
			}
		}
	}

	if testEqSlice2(s, u) {
		fmt.Fprintf(writer, "%v\n", "possible")
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				fmt.Fprintf(writer, "%v", t[i][j])
			}
			fmt.Fprintf(writer, "\n")
		}

	} else {
		fmt.Fprintf(writer, "%v\n", "impossible")
	}
}
