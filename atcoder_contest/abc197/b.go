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
	var x int
	fmt.Fscan(reader, &x)
	var y int
	fmt.Fscan(reader, &y)
	x--
	y--

	var a = make([][]string, h)
	for i := 0; i < h; i++ {
		var ss string
		fmt.Fscan(reader, &ss)
		a[i] = strings.Split(ss, "")
	}

	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}

	ans := 0
	for i := 0; i < 4; i++ {
		nx := x + dx[i]
		ny := y + dy[i]
		for 0 <= nx && nx < h && 0 <= ny && ny < w && a[nx][ny] != "#" {
			ans++
			nx += dx[i]
			ny += dy[i]
		}
	}

	fmt.Fprintf(writer, "%v\n", ans+1)
}
