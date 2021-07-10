package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var tmp string
	fmt.Fscan(reader, &tmp)

	s := strings.Split(tmp, "")
	t := strings.Split("kyoprotenkei90", "")

	sort.Strings(s)
	sort.Strings(t)

	if strings.Join(s, "") == strings.Join(t, "") {
		fmt.Fprintf(writer, "%v\n", "Yes")
	} else {
		fmt.Fprintf(writer, "%v\n", "No")
	}
}
