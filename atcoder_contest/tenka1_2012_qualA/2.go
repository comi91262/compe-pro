package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Scan()
	s := sc.Text()
	r := regexp.MustCompile(` +`)
	fmt.Printf("%v\n", r.ReplaceAllString(s, ","))
}
