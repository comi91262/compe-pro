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
	var a int
	fmt.Fscan(reader, &a)
	var b int
	fmt.Fscan(reader, &b)

	ans := []int{}
	if a == b {
		for a > 0 {
			ans = append(ans, a)
			ans = append(ans, -a)
			a--
		}
	} else {
		if a > b {
			sum := 0
			for a >= b {
				ans = append(ans, a+1)
				sum += a + 1
				a--
			}
			ans = append(ans, -sum)
			for b-1 > 0 {
				ans = append(ans, b)
				ans = append(ans, -b)
				b--
			}
		} else {
			sum := 0
			for a <= b {
				ans = append(ans, -b-1)
				sum += -b - 1
				b--
			}
			ans = append(ans, -sum)
			for a-1 > 0 {
				ans = append(ans, a)
				ans = append(ans, -a)
				a--
			}
		}
	}

	for i := 0; i < len(ans); i++ {
		if i > 0 {
			fmt.Fprintf(writer, " ")
		}
		fmt.Fprintf(writer, "%v", ans[i])
	}
	fmt.Fprintf(writer, "\n")
}
