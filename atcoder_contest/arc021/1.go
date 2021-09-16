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
	a := make([][]int, 4)
	for i := 0; i < 4; i++ {
		a[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			fmt.Fscan(reader, &a[i][j])
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if a[i][j] == a[i][j+1] {
				fmt.Fprintf(writer, "%v\n", "CONTINUE")
				return
			}
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if a[j][i] == a[j+1][i] {
				fmt.Fprintf(writer, "%v\n", "CONTINUE")
				return
			}
		}
	}

	fmt.Fprintf(writer, "%v\n", "GAMEOVER")
}
