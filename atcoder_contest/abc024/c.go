package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var d int
	fmt.Fscan(reader, &d)
	var k int
	fmt.Fscan(reader, &k)

	var l = make([]int, d)
	var r = make([]int, d)
	for i := 0; i < d; i++ {
		fmt.Fscan(reader, &l[i])
		fmt.Fscan(reader, &r[i])
	}

	var s = make([]int, k)
	var t = make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Fscan(reader, &s[i])
		fmt.Fscan(reader, &t[i])
	}

	for i := 0; i < k; i++ {
		st := s[i]
		en := t[i]
		dir := s[i] < t[i]

		ans := 0
		for j := 0; j < d; j++ {
			ans++
			if l[j] <= st && st <= r[j] {
				if l[j] <= en && en <= r[j] {
					break
				} else {
					if dir {
						st = r[j]
					} else {
						st = l[j]
					}
				}
			}
		}
		fmt.Fprintf(writer, "%v\n", ans)
	}

}
