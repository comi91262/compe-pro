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

	var d int
	fmt.Fscan(reader, &d)
	var n int
	fmt.Fscan(reader, &n)

	if d == 0 {
		cnt := 0
		for i := 1; ; i++ {
			if i%100 != 0 {
				cnt++
			}
			if cnt == n {
				fmt.Fprintf(writer, "%v\n", i)
				return
			}
		}
	}

	if d == 1 {
		cnt := 0
		for i := 100; ; i++ {
			if i%10000 == 0 {
				continue
			}
			if i%100 == 0 {
				cnt++
			}
			if cnt == n {
				fmt.Fprintf(writer, "%v\n", i)
				return
			}
		}
	}

	if d == 2 {
		cnt := 0
		for i := 10000; ; i++ {
			if i%1000000 == 0 {
				continue
			}
			if i%10000 == 0 {
				cnt++
			}
			if cnt == n {
				fmt.Fprintf(writer, "%v\n", i)
				return
			}
		}
	}
}
