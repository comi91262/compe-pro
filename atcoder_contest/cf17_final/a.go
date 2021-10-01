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

	var s string
	fmt.Fscan(reader, &s)
	t := []string{"KIHBR", "KIHABR", "AKIHABR", "KIHBRA", "AKIHBRA", "AKIHABARA", "KIHBARA", "AKIHBAR", "KIHABRA", "KIHABARA", "AKIHBR", "KIHBAR", "KIHABAR", "AKIHABAR", "AKIHABRA", "AKIHBARA"}
	for i := range t {
		if t[i] == s {
			fmt.Fprintf(writer, "%v\n", "YES")
			return
		}
	}
	fmt.Fprintf(writer, "%v\n", "NO")

}
