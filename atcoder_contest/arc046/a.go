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
	ans := []int{}

	for i := 1; i < 1000000; i++ {
		if i < 10 {
			ans = append(ans, i)
			continue
		}

		isValid := true
		j := i
		d := i % 10
		for j/10 > 0 {
			j /= 10
			if j%10 != d {
				isValid = false
				break
			}
		}
		if isValid {
			ans = append(ans, i)
		}
	}
	fmt.Fprintf(writer, "%v\n", ans[n-1])

}
