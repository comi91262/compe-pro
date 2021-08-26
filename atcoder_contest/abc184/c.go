package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x * -1
	}
}

func main() {
	defer writer.Flush()

	var r1 int
	fmt.Fscan(reader, &r1)
	var c1 int
	fmt.Fscan(reader, &c1)

	var r2 int
	fmt.Fscan(reader, &r2)
	var c2 int
	fmt.Fscan(reader, &c2)

	if r1 == r2 && c1 == c2 {
		fmt.Fprintf(writer, "%v\n", 0)
		return
	}

	for i := -3; i <= 3; i++ {
		for j := -3; j <= 3; j++ {
			if abs(i)+abs(j) > 4 {
				continue
			}
			if r1+i == r2 && c1+j == c2 {
				fmt.Fprintf(writer, "%v\n", 1)
				return
			}
		}
	}

	dx := r1 - r2
	dy := c1 - c2

	if dy == dx || dy == -dx {
		fmt.Fprintf(writer, "%v\n", 1)
		return
	}

	for i := -3; i <= 3; i++ {
		for j := -3; j <= 3; j++ {
			if abs(i)+abs(j) > 4 {
				continue
			}
			r := r1 + i
			c := c1 + j

			dx := r - r2
			dy := c - c2

			if dy == dx || dy == -dx {
				fmt.Fprintf(writer, "%v\n", 2)
				return
			}
		}
	}

	for i := -6; i <= 6; i++ {
		for j := -6; j <= 6; j++ {
			if abs(i)+abs(j) > 7 {
				continue
			}

			if r1+i == r2 && c1+j == c2 {
				fmt.Fprintf(writer, "%v\n", 2)
				return
			}
		}
	}

	if abs(r1-c1)%2 != abs(r2-c2)%2 {
		fmt.Fprintf(writer, "%v\n", 3)
	} else {
		fmt.Fprintf(writer, "%v\n", 2)
	}
}
