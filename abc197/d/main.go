package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var wtr = bufio.NewWriter(os.Stdout)

func main() {
	var in, ix, iy, ixh, iyh int
	fmt.Scanf("%d", &in)
	fmt.Scanf("%d %d", &ix, &iy)
	fmt.Scanf("%d %d", &ixh, &iyh)

	var n, x, y, xh, yh float64
	n = float64(in)
	x = float64(ix)
	y = float64(iy)
	xh = float64(ixh)
	yh = float64(iyh)
	println(n, x, y, xh, yh)

	var xc, yc, degree, rx, ry float64
	xc = (x + xh) / 2
	yc = (y + yh) / 2
	degree = 360 / n

	rx = math.Cos(degree*math.Pi/180)*(x-xc) - math.Sin(degree*math.Pi/180)*(y-yc) + xc
	ry = math.Sin(degree*math.Pi/180)*(x-xc) + math.Cos(degree*math.Pi/180)*(y-yc) + yc

	fmt.Fprintf(wtr, "%f %f", rx, ry)
	_ = wtr.Flush()
}
