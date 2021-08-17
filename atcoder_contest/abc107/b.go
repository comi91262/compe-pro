package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func contains(x int, a []int) bool {
	for _, e := range a {
		if x == e {
			return true
		}
	}
	return false
}

func main() {
	defer writer.Flush()

	var h int
	fmt.Fscan(reader, &h)
	var w int
	fmt.Fscan(reader, &w)

	a := make([][]string, h)
	for i := 0; i < h; i++ {
		var ss string
		fmt.Fscan(reader, &ss)
		a[i] = strings.Split(ss, "")
	}

	hi := []int{}
	for i := 0; i < h; i++ {
		flag := true
		for j := 0; j < w; j++ {
			if a[i][j] == "#" {
				flag = false
				break
			}
		}
		if flag {
			hi = append(hi, i)
		}
	}

	wi := []int{}
	for i := 0; i < w; i++ {
		flag := true
		for j := 0; j < h; j++ {
			if a[j][i] == "#" {
				flag = false
				break
			}
		}
		if flag {
			wi = append(wi, i)
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if !contains(i, hi) && !contains(j, wi) {
				fmt.Fprintf(writer, "%v", a[i][j])
			}
		}
		if !contains(i, hi) {
			fmt.Fprintf(writer, "\n")
		}
	}

}
