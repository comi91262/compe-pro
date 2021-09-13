package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	s := []string{}
	for i := 1; i <= 1000; i++ {
		s = append(s, strconv.Itoa(i))
	}
	sort.Strings(s)
	for i := range s {
		fmt.Fprintf(writer, "%v\n", s[i])
	}

}
