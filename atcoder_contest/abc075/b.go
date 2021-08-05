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

	a := make([][]string, h)
	for i := 0; i < h; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		a[i] = strings.Split(tmp, "")
	}

	b := make([][]int, h)
	for i := 0; i < h; i++ {
		b[i] = make([]int, w)
		for j := 0; j < w; j++ {
			fmt.Fscan(reader, &b[i][j])
		}
	}
	var dx = [8]int{1, 0, -1, 0, 1, 1, -1, -1}
	var dy = [8]int{0, 1, 0, -1, 1, -1, 1, -1}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if a[i][j] == "#" {
				b[i][j] = -1
				continue
			}
			cnt := 0
			for k := 0; k < 8; k++ {
				x := i + dx[k]
				y := j + dy[k]
				if 0 <= x && x < h && 0 <= y && y < w && a[x][y] == "#" {
					cnt++
				}
			}
			b[i][j] = cnt
		}
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			switch b[i][j] {
			case -1:
				fmt.Fprintf(writer, "%v", "#")
			default:
				fmt.Fprintf(writer, "%v", b[i][j])
			}
		}
		fmt.Fprintf(writer, "\n")
	}
}
