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
	var m int
	fmt.Fscan(reader, &m)

	var a = make([][]string, n)
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		a[i] = strings.Split(tmp, "")
	}

	var b = make([][]string, m)
	for i := 0; i < m; i++ {
		var tmp string
		fmt.Fscan(reader, &tmp)
		b[i] = strings.Split(tmp, "")
	}

	for i := 0; i < n-m+1; i++ {
		for j := 0; j < n-m+1; j++ {
			flag := true
			for k := 0; k < m; k++ {
				for l := 0; l < m; l++ {
					if a[i+k][j+l] != b[k][l] {
						flag = false
					}
				}
			}
			if flag {
				fmt.Fprintf(writer, "%v\n", "Yes")
				return
			}
		}
	}
	fmt.Fprintf(writer, "%v\n", "No")

}

