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
	s := make([][]string, n)
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		s[i] = strings.Split(tmp, "")
	}

	for i := n - 1; i >= 0; i-- {
		for j := 1; j <= 2*(n-1)-1; j++ {
			if s[i][j] == "." || s[i][j] == "X" {
				continue
			}
			if i+1 < n && 0 <= j-1 && s[i+1][j-1] == "X" {
				s[i][j] = "X"
				continue
			}
			if i+1 < n && s[i+1][j] == "X" {
				s[i][j] = "X"
				continue
			}
			if i+1 < n && j+1 < 2*n-1 && s[i+1][j+1] == "X" {
				s[i][j] = "X"
				continue
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Fprintf(writer, "%v\n", strings.Join(s[i], ""))
	}

}
