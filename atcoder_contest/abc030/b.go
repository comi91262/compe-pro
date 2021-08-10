package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func min(arg ...float64) float64 {
	min := arg[0]
	for _, x := range arg {
		if min > x {
			min = x
		}
	}
	return min
}
func main() {

	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)
	var m int
	fmt.Fscan(reader, &m)

	nDeg := float64((90-n*30)%360) - float64(m)/2.0
	mDeg := (90 - m*6) % 360

	diff := math.Abs(float64(mDeg) - nDeg)
	fmt.Fprintf(writer, "%v\n", min(diff, 360-diff))

}
