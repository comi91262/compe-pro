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

	var deg, dis int
	fmt.Fscan(reader, &deg, &dis)

	r := ""
	switch {
	case 113 <= deg && deg <= 337:
		r += "NNE"
	case 338 <= deg && deg <= 562:
		r += "NE"
	case 563 <= deg && deg <= 787:
		r += "ENE"
	case 788 <= deg && deg <= 1012:
		r += "E"
	case 1013 <= deg && deg <= 1237:
		r += "ESE"
	case 1238 <= deg && deg <= 1462:
		r += "SE"
	case 1463 <= deg && deg <= 1687:
		r += "SSE"
	case 1688 <= deg && deg <= 1912:
		r += "S"
	case 1913 <= deg && deg <= 2137:
		r += "SSW"
	case 2138 <= deg && deg <= 2362:
		r += "SW"
	case 2363 <= deg && deg <= 2587:
		r += "WSW"
	case 2588 <= deg && deg <= 2812:
		r += "W"
	case 2813 <= deg && deg <= 3037:
		r += "WNW"
	case 3038 <= deg && deg <= 3262:
		r += "NW"
	case 3263 <= deg && deg <= 3487:
		r += "NNW"
	default:
		r += "N"
	}

	r2 := 0
	dis *= 10
	switch {
	case 0 <= dis && dis <= (20+4)*6:
		r = "C"
		r2 = 0
	case (30-5)*6 <= dis && dis <= (150+4)*6:
		r2 = 1
	case (160-5)*6 <= dis && dis <= (330+4)*6:
		r2 = 2
	case (340-5)*6 <= dis && dis <= (540+4)*6:
		r2 = 3
	case (550-5)*6 <= dis && dis <= (790+4)*6:
		r2 = 4
	case (800-5)*6 <= dis && dis <= (1070+4)*6:
		r2 = 5
	case (1080-5)*6 <= dis && dis <= (1380+4)*6:
		r2 = 6
	case (1390-5)*6 <= dis && dis <= (1710+4)*6:
		r2 = 7
	case (1720-5)*6 <= dis && dis <= (2070+4)*6:
		r2 = 8
	case (2080-5)*6 <= dis && dis <= (2440+4)*6:
		r2 = 9
	case (2450-5)*6 <= dis && dis <= (2840+4)*6:
		r2 = 10
	case (2850-5)*6 <= dis && dis <= (3260+4)*6:
		r2 = 11
	case (3270-5)*6 <= dis:
		r2 = 12
	}

	fmt.Fprintf(writer, "%v %v\n", r, r2)

}
