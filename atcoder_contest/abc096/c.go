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

	var dx = [4]int{1, 0, -1, 0}
	var dy = [4]int{0, 1, 0, -1}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == "." {
				continue
			}
			flag := false
			for k := 0; k < 4; k++ {
				nx := i + dx[k]
				ny := j + dy[k]
				if nx >= h || nx < 0 || ny >= w || ny < 0 {
					continue
				}
				if s[nx][ny] == "#" {
					flag = true
				}
			}
			if !flag {
				fmt.Fprintf(writer, "%v\n", "No")
				return
			}

		}
	}

	fmt.Fprintf(writer, "%v\n", "Yes")
}
 