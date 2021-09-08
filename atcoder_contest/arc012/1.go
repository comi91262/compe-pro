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
	var d string
	fmt.Fscan(reader, &d)

	switch d {
	case "Sunday":
		fmt.Fprintf(writer, "%v\n", 0)
	case "Monday":
		fmt.Fprintf(writer, "%v\n", 5)
	case "Tuesday":
		fmt.Fprintf(writer, "%v\n", 4)
	case "Wednesday":
		fmt.Fprintf(writer, "%v\n", 3)
	case "Thursday":
		fmt.Fprintf(writer, "%v\n", 2)
	case "Friday":
		fmt.Fprintf(writer, "%v\n", 1)
	case "Saturday":
		fmt.Fprintf(writer, "%v\n", 0)
	}

}
