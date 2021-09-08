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

	s := make([][]string, h)
	for i := 0; i < h; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		s[i] = strings.Split(tmp, "")
	}

	cnt := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w-1; j++ {
			if s[i][j] == "." && s[i][j+1] == "." {
				cnt++
			}
		}
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h-1; j++ {
			if s[j][i] == "." && s[j+1][i] == "." {
				cnt++
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", cnt)
}
